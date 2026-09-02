// Package parsepair — Gopher Workplace challenge.
package parsepair

import (
	"errors"
	"strings"
)

// Parse failures.
var (
	ErrNoSeparator = errors.New("missing =")
	ErrEmptyKey    = errors.New("empty key")
)

// ParsePair splits s into key and value at the first "=".
//
// Examples:
//
//	ParsePair("HOST=local") => "HOST", "local", nil
//	ParsePair("HOST")       => "", "", ErrNoSeparator
func ParsePair(s string) (string, string, error) {
	// TODO(candidate): implement this.
	_ = strings.Cut
	panic("not implemented")
}
