// Package joinflatten — Gopher Workplace challenge.
package joinflatten

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
	ErrC = errors.New("c")
)

// Leaves returns every leaf error inside err, depth first.
//
// Examples:
//
//	Leaves(nil) => nil
func Leaves(err error) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
