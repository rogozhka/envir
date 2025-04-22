package envir

import (
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/rogozhka/envir/decoders"
)

var DefaultDecoder = decoders.NewEnv()

type lookupFile struct {
	path    string
	decoder decoderInterface

	cutPrefix string

	mustDecode bool

	mx        sync.RWMutex
	isDecoded bool
	mm        map[string]any
}

func NewLookupFile(path string, opts ...func(*lookupFile)) *lookupFile {
	p := &lookupFile{
		path: path,
		mm:   map[string]any{},
	}

	for _, opt := range opts {
		opt(p)
	}

	if p.decoder == nil {
		p.decoder = DefaultDecoder
	}

	return p
}

func WithDecoder(decoder decoderInterface) func(*lookupFile) {
	return func(p *lookupFile) {
		p.decoder = decoder
	}
}

func WithCutPrefix(cutPrefix string) func(*lookupFile) {
	return func(p *lookupFile) {
		prefix := strings.TrimSpace(cutPrefix)
		if prefix[len(prefix)-1] != '_' {
			prefix += "_"
		}

		p.cutPrefix = prefix
	}
}

func WithMustDecode() func(*lookupFile) {
	return func(p *lookupFile) {
		p.mustDecode = true
	}
}

func (p *lookupFile) prepareName(name string) string {
	if name == "" || p.cutPrefix == "" || len(name) <= len(p.cutPrefix) {
		return name
	}
	lp := len(p.cutPrefix)

	if name[:lp] != p.cutPrefix {
		return name
	}

	return name[lp:]
}

func (p *lookupFile) LookupEnv(name string) (string, bool) {
	p.mx.RLock()
	isDecoded := p.isDecoded
	p.mx.RUnlock()

	if !isDecoded {
		if err := p.decodeFile(); err != nil {
			if p.mustDecode {
				panic(fmt.Errorf("decode file | %v", err))
			}
		}
		p.mx.Lock()
		p.isDecoded = true
		p.mx.Unlock()
	}

	p.mx.RLock()
	defer p.mx.RUnlock()

	key := p.prepareName(name)

	raw, ok := p.mm[key]
	if !ok {
		return "", false
	}

	v, ok := raw.(string)
	if !ok {
		panic(fmt.Errorf("cannot convert to string | %s | %v", key, raw))
	}

	return v, true
}

func (p *lookupFile) decodeFile() error {
	p.mx.Lock()
	defer p.mx.Unlock()

	bb, err := os.ReadFile(p.path)
	if err != nil {
		return fmt.Errorf("read file | %v | %w", p.path, err)
	}

	if err = p.decoder.Decode(bb, p.mm); err != nil {
		return fmt.Errorf("decoder | %w", err)
	}

	return nil
}
