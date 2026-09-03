// Package validationerr — Gopher Workplace challenge.
package validationerr

import "fmt"

// ValidationError reports a single rejected field.
type ValidationError struct {
	Field  string
	Reason string
}

// Error implements the error interface as "<Field>: <Reason>".
func (e *ValidationError) Error() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}

// NewValidation builds a *ValidationError for the given field.
//
// Examples:
//
//	NewValidation("email", "is required") => "email: is required"
func NewValidation(field, reason string) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
