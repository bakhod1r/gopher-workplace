// Package findsentinel — Gopher Workplace challenge.
package findsentinel

import "errors"

// Replica failures used by the tests.
var (
	ErrTimeout = errors.New("timeout")
	ErrOther   = errors.New("other failure")
)

// FirstMatch returns the first error in errs matching target.
//
// Examples:
//
//	FirstMatch([]error{ErrOther, ErrTimeout}, ErrTimeout) => ErrTimeout
//	FirstMatch([]error{ErrOther}, ErrTimeout)             => nil
func FirstMatch(errs []error, target error) error {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
