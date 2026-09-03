// Package firstleaf — Gopher Workplace challenge.
package firstleaf

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Origin returns the leftmost leaf of err's tree.
//
// Examples:
//
//	Origin(nil) => nil
func Origin(err error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
