package decoders

import "errors"

var (
	ErrArguments   = errors.New("arguments")
	ErrNoSuchEntry = errors.New("no such entry")
)
