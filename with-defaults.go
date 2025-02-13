package envir

import "time"

// String read env as string w/ defaults if no such env var.
func String(name, defaultValue string) string {
	if !IsSet(name) {
		return defaultValue
	}

	return MustString(name)
}

// Int reads env as integer w/ defaults if no such env var.
func Int(name string, defaultValue int) int {
	if !IsSet(name) {
		return defaultValue
	}

	return MustInt(name)
}

func Int64(name string, defaultValue int64) int64 {
	if !IsSet(name) {
		return defaultValue
	}

	return MustInt64(name)
}

// Uint read env as unsigned integer w/ defaults if no such env var.
func Uint(name string, defaultValue uint) uint {
	if !IsSet(name) {
		return defaultValue
	}

	return MustUint(name)
}

func Uint64(name string, defaultValue uint64) uint64 {
	if !IsSet(name) {
		return defaultValue
	}

	return MustUint64(name)
}

func Float64(name string, defaultValue float64) float64 {
	if !IsSet(name) {
		return defaultValue
	}

	return MustFloat64(name)
}

// Bool read env as unsigned integer w/ defaults if no such env var.
func Bool(name string, defaultValue bool) bool {
	if !IsSet(name) {
		return defaultValue
	}

	return MustBool(name)
}

// Duration read env as time duration w/ defaults if no such env var.
func Duration(name string, defaultValue time.Duration) time.Duration {
	if !IsSet(name) {
		return defaultValue
	}

	return MustDuration(name)
}
