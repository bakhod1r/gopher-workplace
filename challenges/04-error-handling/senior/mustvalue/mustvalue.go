// Package mustvalue — Gopher Workplace challenge.
package mustvalue

import "errors"

// ErrLoad is a stand-in load failure used by the tests.
var ErrLoad = errors.New("load failed")

// Must returns v, panicking when err is non-nil.
//
// Examples:
//
//	Must(42, nil)     => 42
//	Must(0, ErrLoad)  => panics with ErrLoad
func Must(v int, err error) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
