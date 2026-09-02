// Package counterrs — Gopher Workplace challenge.
package counterrs

import "errors"

// ErrX is a stand-in failure used by the tests.
var ErrX = errors.New("record failed")

// CountErrors returns how many entries of errs are non-nil.
//
// Examples:
//
//	CountErrors([]error{nil, ErrX, ErrX}) => 2
//	CountErrors(nil)                      => 0
func CountErrors(errs []error) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
