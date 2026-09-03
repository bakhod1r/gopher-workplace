// Package leafsummary — Gopher Workplace challenge.
package leafsummary

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Summary counts leaf messages inside err.
//
// Examples:
//
//	Summary(nil) => {}
func Summary(err error) map[string]int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
