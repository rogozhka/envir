package envir

//go:generate docker run -v ${PWD}:/w rogozhka/go-generate-mockgen:namefix -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type decoderInterface interface {
	Decode([]byte, map[string]any) error
}
