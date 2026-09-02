// Package haserror — Gopher Workplace challenge.
package haserror

import "errors"

// ErrSample is a stand-in failure used by the tests.
var ErrSample = errors.New("sample failure")

// HasError reports whether err represents a failure.
//
// Examples:
//
//	HasError(nil)       => false
//	HasError(ErrSample) => true
func HasError(err error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
