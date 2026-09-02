// Package firstfail — Gopher Workplace challenge.
package firstfail

import "errors"

// ErrNoFailure reports that no entry failed.
var ErrNoFailure = errors.New("no failure found")

// ErrStep is a stand-in failure used by the tests.
var ErrStep = errors.New("step failed")

// FirstFail returns the index of the first non-nil error in errs.
//
// Examples:
//
//	FirstFail([]error{nil, ErrStep}) => 1, nil
//	FirstFail([]error{nil, nil})     => -1, ErrNoFailure
func FirstFail(errs []error) (int, error) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
