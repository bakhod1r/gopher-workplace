// Package chainto — Gopher Workplace challenge.
package chainto

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Above returns the messages layered above target in err's chain.
//
// Examples:
//
//	Above(ErrA, ErrA) => []
func Above(err, target error) []string {
	// TODO(candidate): implement this.
	_ = errors.Unwrap
	panic("not implemented")
}
