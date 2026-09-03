// Package depthlimit — Gopher Workplace challenge.
package depthlimit

import "errors"

// ErrA is a stand-in failure used by the tests.
var ErrA = errors.New("a")

// Within reports whether target appears within max links of err.
//
// Examples:
//
//	Within(ErrA, ErrA, 1) => true
func Within(err, target error, max int) bool {
	// TODO(candidate): implement this.
	_ = errors.Unwrap
	panic("not implemented")
}
