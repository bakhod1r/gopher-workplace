// Package stripwrap — Gopher Workplace challenge.
package stripwrap

import "errors"

// Public failures used by the tests.
var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("denied")
)

// Strip collapses a known cause to its sentinel, leaving others untouched.
//
// Examples:
//
//	Strip(fmt.Errorf("a: %w", ErrNotFound)) => ErrNotFound
//	Strip(nil)                              => nil
func Strip(err error) error {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
