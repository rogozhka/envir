package envir

import (
	"fmt"
	"strings"
	"time"
)

var (
	DefaultLookup   = NewLookupOs()
	DefaultExpander = asisExpander{}
)

type oinment struct {
	prefix   string
	lookuper lookupInterface
	expander expanderInterface
}

func New(options ...func(oinment *oinment)) *oinment {
	p := &oinment{}

	for _, option := range options {
		option(p)
	}

	if p.lookuper == nil {
		p.lookuper = DefaultLookup
	}
	if p.expander == nil {
		p.expander = DefaultExpander
	}

	return p
}

func WithPrefix(prefix string) func(*oinment) {
	return func(p *oinment) {
		p.SetPrefix(prefix)
	}
}

func WithLookup(lookuper lookupInterface) func(*oinment) {
	return func(p *oinment) {
		p.lookuper = lookuper
	}
}

func WithEnvSubst() func(*oinment) {
	return func(p *oinment) {
		p.expander = &envsubstExpander{}
	}
}

func (p *oinment) SetPrefix(prefix string) {
	p.prefix = strings.TrimSpace(prefix)
	if p.prefix == "" {
		return
	}

	if p.prefix[len(p.prefix)-1] != '_' {
		p.prefix += "_"
	}
}

func (p *oinment) Key(name string) string {
	return p.prefix + name
}

func (p *oinment) Value(unprefixed string) string {
	name := p.Key(unprefixed)

	v, _ := p.lookuper.LookupEnv(name)

	expanded, err := p.expander.Expand(v)
	if err != nil {
		return ""
	}

	return expanded
}

func (p *oinment) IsSet(unprefixed string) bool {
	name := p.Key(unprefixed)

	_, ok := p.lookuper.LookupEnv(name)
	return ok
}

func (p *oinment) MustString(unprefixed string) string {
	name := p.Key(unprefixed)

	tr := strings.TrimSpace(name)
	v, ok := p.lookuper.LookupEnv(tr)
	if !ok {
		panic(fmt.Errorf("cannot find env | %s", tr))
	}

	expanded, err := p.expander.Expand(v)
	if err != nil {
		panic(fmt.Errorf("expand | %v | %v", v, err))
	}

	return expanded
}

func (p *oinment) MustInt(unprefixed string) int {
	return intFromString(p.MustString(unprefixed))
}

func (p *oinment) MustInt64(unprefixed string) int64 {
	return int64FromString(p.MustString(unprefixed))
}

func (p *oinment) MustUint(unprefixed string) uint {
	return uint(intFromString(p.MustString(unprefixed)))
}

func (p *oinment) MustUint64(unprefixed string) uint64 {
	return uint64FromString(p.MustString(unprefixed))
}

func (p *oinment) MustFloat64(unprefixed string) float64 {
	return float64FromString(p.MustString(unprefixed))
}

func (p *oinment) MustBool(unprefixed string) bool {
	return boolFromString(p.MustString(unprefixed))
}

func (p *oinment) MustDuration(unprefixed string) time.Duration {
	return durationFromString(p.MustString(unprefixed))
}

func (p *oinment) String(unprefixed, defaultValue string) string {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustString(unprefixed)
}

func (p *oinment) Int(unprefixed string, defaultValue int) int {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustInt(unprefixed)
}

func (p *oinment) Int64(unprefixed string, defaultValue int64) int64 {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustInt64(unprefixed)
}

func (p *oinment) Uint(unprefixed string, defaultValue uint) uint {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustUint(unprefixed)
}

func (p *oinment) Uint64(unprefixed string, defaultValue uint64) uint64 {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustUint64(unprefixed)
}

func (p *oinment) Float64(unprefixed string, defaultValue float64) float64 {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustFloat64(unprefixed)
}

func (p *oinment) Bool(unprefixed string, defaultValue bool) bool {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustBool(unprefixed)
}

func (p *oinment) Duration(unprefixed string, defaultValue time.Duration) time.Duration {
	if !p.IsSet(unprefixed) {
		return defaultValue
	}

	return p.MustDuration(unprefixed)
}
