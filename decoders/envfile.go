package decoders

import (
	"bytes"
	"fmt"

	"github.com/subosito/gotenv"
)

type envDecoder struct{}

func NewEnv() *envDecoder {
	return &envDecoder{}
}

func (p *envDecoder) Decode(bb []byte, mm map[string]any) error {
	if bb == nil || mm == nil {
		return ErrArguments
	}

	r := bytes.NewReader(bb)

	values, err := gotenv.StrictParse(r)
	if err != nil {
		return fmt.Errorf("parse | %w", err)
	}

	for key, value := range values {
		mm[key] = value
	}

	return nil
}
