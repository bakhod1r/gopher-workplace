// Package lasterr — Gopher Workplace challenge.
package lasterr

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("attempt a failed")
	ErrB = errors.New("attempt b failed")
)

// LastError returns the last non-nil error in errs, or nil.
//
// Examples:
//
//	LastError([]error{ErrA, nil, ErrB}) => ErrB
//	LastError([]error{ErrA, nil})       => ErrA
func LastError(errs []error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
