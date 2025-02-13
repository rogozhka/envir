package envir

//go:generate docker run -v ${PWD}:/w rogozhka/go-generate-mockgen:namefix -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type expanderInterface interface {
	Expand(value string) (string, error)
}

type asisExpander struct{}

func (e asisExpander) Expand(value string) (string, error) {
	return value, nil
}
