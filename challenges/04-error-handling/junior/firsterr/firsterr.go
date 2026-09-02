// Package firsterr — Gopher Workplace challenge.
package firsterr

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("check a failed")
	ErrB = errors.New("check b failed")
	ErrC = errors.New("check c failed")
)

// FirstError returns the first non-nil error in errs, or nil.
//
// Examples:
//
//	FirstError([]error{nil, ErrB}) => ErrB
//	FirstError([]error{nil, nil})  => nil
func FirstError(errs []error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
