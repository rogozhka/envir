package envir

//go:generate docker run -v ${PWD}:/w rogozhka/go-generate-mockgen:namefix -line=$GOLINE -source=$GOFILE -package=$GOPACKAGE
type lookupInterface interface {
	LookupEnv(key string) (string, bool)
}
