package envir

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func intFromString(s string) int {
	res, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Errorf("cannot convert string to int | %s | %v", s, err))
	}

	return res
}

func int64FromString(s string) int64 {
	res, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("cannot convert string to int64 | %s | %v", s, err))
	}

	return res
}

func uint64FromString(s string) uint64 {
	res, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		panic(fmt.Errorf("cannot convert string to uint64 | %s | %v", s, err))
	}

	return res
}

func float64FromString(s string) float64 {
	res, err := strconv.ParseFloat(s, 10)
	if err != nil {
		panic(fmt.Errorf("cannot convert string to float64 | %s | %v", s, err))
	}

	return res
}

func boolFromString(str string) bool {
	s := strings.ToLower(str)
	switch s {
	case "1", "yes", "y", "set", "true", "t", "checked", "da", "ok", "ок", "д", "да":
		return true
	}

	return false
}

func durationFromString(str string) time.Duration {
	d, err := time.ParseDuration(str)
	if err != nil {
		return 0
	}

	return d
}
