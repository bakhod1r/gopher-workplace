// Package rootcause — Gopher Workplace challenge.
package rootcause

import "errors"

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base failure")

// Root returns the deepest error in err's chain.
//
// Examples:
//
//	Root(fmt.Errorf("a: %w", ErrBase)) => ErrBase
//	Root(ErrBase)                      => ErrBase
func Root(err error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
