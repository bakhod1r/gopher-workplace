// Package asall — Gopher Workplace challenge.
package asall

import "errors"

// ErrOther is a stand-in unrelated failure used by the tests.
var ErrOther = errors.New("other")

// FieldError reports a failure on a named field.
type FieldError struct {
	Name string
}

// Error implements the error interface.
func (e *FieldError) Error() string {
	return "field " + e.Name
}

// All returns every *FieldError inside err.
//
// Examples:
//
//	All(nil) => nil
func All(err error) []*FieldError {
	// TODO(candidate): implement this.
	panic("not implemented")
}
