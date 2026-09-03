// Package asinjoin — Gopher Workplace challenge.
package asinjoin

import (
	"errors"
	"fmt"
)

// ErrOther is a stand-in non-field failure used by the tests.
var ErrOther = errors.New("other failure")

// FieldError reports a failure attached to a named field.
type FieldError struct {
	Name string
}

// Error implements the error interface.
func (e *FieldError) Error() string {
	return fmt.Sprintf("field %s is invalid", e.Name)
}

// FirstField returns the first *FieldError inside err.
//
// Examples:
//
//	FirstField(&FieldError{Name: "age"}) => the *FieldError, true
//	FirstField(ErrOther)                 => nil, false
func FirstField(err error) (*FieldError, bool) {
	// TODO(candidate): implement this.
	_ = errors.As
	panic("not implemented")
}
