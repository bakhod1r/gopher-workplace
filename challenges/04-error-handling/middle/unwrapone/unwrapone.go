// Package unwrapone — Gopher Workplace challenge.
package unwrapone

import "errors"

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base failure")

// Cause returns the error directly wrapped by err, or nil.
//
// Examples:
//
//	Cause(fmt.Errorf("a: %w", ErrBase)) => ErrBase
//	Cause(ErrBase)                      => nil
func Cause(err error) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
