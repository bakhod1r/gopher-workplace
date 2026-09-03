// Package wrapdepth — Gopher Workplace challenge.
package wrapdepth

import "errors"

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base failure")

// Depth returns how many errors are in err's chain.
//
// Examples:
//
//	Depth(nil)     => 0
//	Depth(ErrBase) => 1
func Depth(err error) int {
	// TODO(candidate): implement this.
	_ = errors.Unwrap
	panic("not implemented")
}
