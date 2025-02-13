package envir

type lookupComposition struct {
	primary   lookupInterface
	optionals []lookupInterface
}

func NewLookupComposition(primary lookupInterface, opts ...func(*lookupComposition)) *lookupComposition {
	p := &lookupComposition{
		primary: primary,
	}

	for _, opt := range opts {
		opt(p)
	}

	return p
}

func WithOptional(optionals ...lookupInterface) func(*lookupComposition) {
	return func(p *lookupComposition) {
		for _, opt := range optionals {
			p.optionals = append(p.optionals, opt)
		}
	}
}

// LookupEnv checks primary then optionals.
func (p *lookupComposition) LookupEnv(name string) (string, bool) {
	v, ok := p.primary.LookupEnv(name)
	if ok {
		return v, true
	}

	for _, opt := range p.optionals {
		v, ok := opt.LookupEnv(name)
		if ok {
			return v, true
		}
	}

	return "", false
}
