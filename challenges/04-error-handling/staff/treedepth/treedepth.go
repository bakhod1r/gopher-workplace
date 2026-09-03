// Package treedepth — Gopher Workplace challenge.
package treedepth

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Depth returns the height of err's tree.
//
// Examples:
//
//	Depth(nil)     => 0
//	Depth(ErrA)    => 1
func Depth(err error) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
