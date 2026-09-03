// Package joinerrs — Gopher Workplace challenge.
package joinerrs

import "errors"

// Rule failures used by the tests.
var (
	ErrA = errors.New("rule a failed")
	ErrB = errors.New("rule b failed")
)

// JoinAll combines every non-nil error in errs into one.
//
// Examples:
//
//	JoinAll([]error{ErrA, ErrB}) => error matching both
//	JoinAll(nil)                 => nil
func JoinAll(errs []error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
