// Package nilsafejoin — Gopher Workplace challenge.
package nilsafejoin

import "errors"

// OpError reports a failed operation.
type OpError struct {
	Op string
}

// Error implements the error interface.
func (e *OpError) Error() string {
	return e.Op + " failed"
}

// Clean joins the entries, skipping nil and typed-nil *OpError values.
//
// Examples:
//
//	Clean() => nil
func Clean(errs ...error) error {
	// TODO(candidate): implement this.
	_ = errors.Join
	panic("not implemented")
}
