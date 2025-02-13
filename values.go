// Package envir reads environment variables.
package envir

import "time"

var env = New()

// IsPresent indicates env var is set (can be empty value).
// DEPRECATED: use IsSet
func IsPresent(name string) bool {
	return IsSet(name)
}

// StringValue returns env variable as string,
// panics in case of error.
// DEPRECATED: use MustString instead.
func StringValue(name string) string {
	return MustString(name)
}

// IntValue returns env variable as integer,
// panics in case of error.
// DEPRECATED: use MustInt.
func IntValue(name string) int {
	return MustInt(name)
}

// UintValue returns env variable as integer,
// panics in case of error.
// DEPRECATED: use MustUint.
func UintValue(name string) uint {
	return MustUint(name)
}

// BoolValue treats env string as boolean: 1, yes, Y, true, etc are true, exceptions are false.
// DEPRECATED: use MustBool instead.
func BoolValue(name string) bool {
	return MustBool(name)
}

// IsSet indicates env var is set (can be empty value).
func IsSet(name string) bool {
	return env.IsSet(name)
}

// Value returns value or empty if not present.
func Value(name string) string {
	return env.Value(name)
}

// MustString returns value or panics if there is no variable present.
func MustString(name string) string {
	return env.MustString(name)
}

// MustInt returns value or panics if there is no variable present.
func MustInt(name string) int {
	return env.MustInt(name)
}

// MustInt64 returns value or panics if there is no variable present.
func MustInt64(name string) int64 {
	return env.MustInt64(name)
}

// MustUint returns value or panics if there is no variable present.
func MustUint(name string) uint {
	return env.MustUint(name)
}

// MustUint64 returns value or panics if there is no variable present.
func MustUint64(name string) uint64 {
	return env.MustUint64(name)
}

// MustFloat64 returns value or panics if there is no variable present.
func MustFloat64(name string) float64 {
	return env.MustFloat64(name)
}

// MustBool returns value or panics if there is no variable present.
func MustBool(name string) bool {
	return env.MustBool(name)
}

// MustDuration returns value or panics if there is no variable present.
func MustDuration(name string) time.Duration {
	return env.MustDuration(name)
}
