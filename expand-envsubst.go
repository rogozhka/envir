package envir

import "github.com/a8m/envsubst"

type envsubstExpander struct{}

func (*envsubstExpander) Expand(s string) (string, error) {
	return envsubst.String(s)
}
