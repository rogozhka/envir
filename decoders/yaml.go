package decoders

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type yamlDecoder struct {
	path string
}

func NewYaml(opts ...func(decoder *yamlDecoder)) *yamlDecoder {
	p := &yamlDecoder{}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

// WithEntriesPath defines path within entries.
// i.e. "env.variables" checks env: variables: and extracts only variables: children;
// dot is a separator.
func WithEntriesPath(path string) func(*yamlDecoder) {
	return func(p *yamlDecoder) {
		p.path = path
	}
}

func (p *yamlDecoder) Decode(bb []byte, mm map[string]any) error {
	if bb == nil || mm == nil {
		return ErrArguments
	}

	extracted := make(map[string]interface{})
	if err := yaml.Unmarshal(bb, &extracted); err != nil {
		return fmt.Errorf("yaml unmarshal | %w", err)
	}

	if err := p.extractIfExists(extracted, mm); err != nil {
		return fmt.Errorf("path extract | %w", err)
	}

	return nil
}

func (p *yamlDecoder) extractIfExists(from, to map[string]any) error {
	if from == nil || to == nil {
		return ErrArguments
	}

	entries := strings.Split(p.path, ".")
	mm := from
	for _, entry := range entries {
		raw, ok := mm[entry]
		if !ok {
			return ErrNoSuchEntry
		}
		mm, ok = raw.(map[string]any)
		if !ok {
			return ErrNoSuchEntry
		}
	}

	for k, v := range mm {
		to[k] = v
	}

	return nil
}
