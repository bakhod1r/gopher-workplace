// Package errorsis — Gopher Workplace challenge.
package errorsis

import "errors"

// ErrNotFound is returned when a key is absent.
var ErrNotFound = errors.New("not found")

// Fetch reads a key, wrapping ErrNotFound with context when it is missing.
//
// Examples:
//
//	Fetch(map[string]string{}, "k") => error "fetch k: not found"
//	Fetch(map[string]string{"k": "v"}, "k") => "v", nil
func Fetch(data map[string]string, key string) (string, error) {
	// TODO(candidate): wrap ErrNotFound with the key.
	panic("not implemented")
}

// IsMissing reports whether err is or wraps ErrNotFound.
func IsMissing(err error) bool {
	// TODO(candidate): inspect the whole chain.
	panic("not implemented")
}

// FetchAll reads every key, returning the first error.
func FetchAll(data map[string]string, keys []string) ([]string, error) {
	// TODO(candidate): collect values, propagate the first failure.
	panic("not implemented")
}
