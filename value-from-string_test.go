package envir

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoolFromString(t *testing.T) {
	is := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{"empty string", "", false},
		{"false cases", "no", false},
		{"false cases", "n", false},
		{"false cases", "unset", false},
		{"false cases", "false", false},
		{"false cases", "f", false},
		{"false cases", "unchecked", false},
		{"false cases", "nein", false},
		{"false cases", "нет", false},
		{"true cases", "1", true},
		{"true cases", "yes", true},
		{"true cases", "y", true},
		{"true cases", "Y", true},
		{"true cases", "set", true},
		{"true cases", "true", true},
		{"true cases", "TRUE", true},
		{"true cases", "t", true},
		{"true cases", "checked", true},
		{"true cases", "da", true},
		{"true cases", "ok", true},
		{"true cases", "ок", true},
		{"true cases", "д", true},
		{"true cases", "да", true},
	}

	t.Parallel()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := boolFromString(tt.input)
			is.Equal(tt.expected, result, "Input: %s", tt.input)
		})
	}
}

func TestFloat64FromString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected float64
	}{
		{"valid float", "123.45", 123.45},
		{"valid float with leading zero", "0.12345", 0.12345},
		{"valid integer as float", "42", 42.0},
		{"negative float", "-123.45", -123.45},
		{"scientific notation", "1e-6", 1e-6},
		{"scientific notation with positive exponent", "1e6", 1e6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := float64FromString(tt.input)
			is.Equal(tt.expected, result, "Input: %s", tt.input)
		})
	}

	// Тесты на неверные входные данные
	invalidTests := []struct {
		name  string
		input string
	}{
		{"invalid input", "abc"},
		{"empty string", ""},
		{"only dot", "."},
		{"multiple dots", "123.45.67"},
	}

	for _, tt := range invalidTests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for input: %s", tt.input)
				}
			}()
			float64FromString(tt.input)
		})
	}
}

func TestUint64FromString(t *testing.T) {
	t.Parallel()
	is := assert.New(t)

	tests := []struct {
		name     string
		input    string
		expected uint64
	}{
		{"valid uint64", "12345", 12345},
		{"zero value", "0", 0},
		{"max uint64", "18446744073709551615", 18446744073709551615},
		{"large number", "9223372036854775807", 9223372036854775807},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := uint64FromString(tt.input)
			is.Equal(tt.expected, result, "Input: %s", tt.input)
		})
	}

	// Тесты на неверные входные данные
	invalidTests := []struct {
		name  string
		input string
	}{
		{"negative number", "-123"},
		{"non-numeric input", "abc"},
		{"empty string", ""},
		{"overflow", "18446744073709551616"},
		{"floating point", "123.45"},
	}

	for _, tt := range invalidTests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("Expected panic for input: %s", tt.input)
				}
			}()
			uint64FromString(tt.input)
		})
	}
}
