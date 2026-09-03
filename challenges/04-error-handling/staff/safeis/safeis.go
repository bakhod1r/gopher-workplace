// Package safeis — Gopher Workplace challenge.
package safeis

import "errors"

// ErrA is a stand-in failure used by the tests.
var ErrA = errors.New("a")

// SafeIs reports whether err matches target, returning false on a panic.
//
// Examples:
//
//	SafeIs(ErrA, ErrA) => true
//	SafeIs(nil, ErrA)  => false
func SafeIs(err, target error) (match bool) {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
