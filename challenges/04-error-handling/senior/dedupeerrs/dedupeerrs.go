// Package dedupeerrs — Gopher Workplace challenge.
package dedupeerrs

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("shard unreachable")
	ErrB = errors.New("shard busy")
)

// Unique returns one error per distinct message, in first-seen order.
//
// Examples:
//
//	Unique([]error{ErrA, ErrA, ErrB}) => [ErrA ErrB]
//	Unique(nil)                       => nil
func Unique(errs []error) []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
