// Package sentineltable — Gopher Workplace challenge.
package sentineltable

import "errors"

// Domain failures used by the tests.
var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("denied")
)

// Rule maps a sentinel onto a status code.
type Rule struct {
	Err  error
	Code int
}

// Status returns the code for the first matching rule.
//
// Examples:
//
//	Status(nil, nil) => 200
func Status(err error, table []Rule) int {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
