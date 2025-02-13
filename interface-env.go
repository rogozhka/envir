package envir

import "time"

var _ envInterface = (*oinment)(nil)

type envInterface interface {
	IsSet(unprefixed string) bool

	MustString(unprefixed string) string

	MustInt(unprefixed string) int

	MustInt64(unprefixed string) int64

	MustUint(unprefixed string) uint

	MustUint64(unprefixed string) uint64

	MustFloat64(unprefixed string) float64

	MustBool(unprefixed string) bool

	String(unprefixed, defaultValue string) string

	Int(unprefixed string, defaultValue int) int

	Int64(unprefixed string, defaultValue int64) int64

	Uint(unprefixed string, defaultValue uint) uint

	Uint64(unprefixed string, defaultValue uint64) uint64

	Float64(unprefixed string, defaultValue float64) float64

	Bool(unprefixed string, defaultValue bool) bool

	Duration(unprefixed string, defaultValue time.Duration) time.Duration
}
