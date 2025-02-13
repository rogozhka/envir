package envir

import "os"

type wrapOs struct {
}

func (wrapOs) LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

func NewLookupOs() wrapOs {
	return wrapOs{}
}
