// Package erriface — Gopher Workplace challenge.
package erriface

// ValidationError reports a bad field.
type ValidationError struct {
	Field   string
	Message string
}

// Error renders the failure.
//
// Examples:
//
//	(&ValidationError{Field: "name", Message: "required"}).Error() => "name: required"
func (e *ValidationError) Error() string {
	// TODO(candidate): "<Field>: <Message>".
	panic("not implemented")
}

// Validate checks a user name.
//
// Examples:
//
//	Validate("ann") => nil
//	Validate("")    => &ValidationError{Field: "name", Message: "required"}
func Validate(name string) error {
	// TODO(candidate): empty name is invalid.
	panic("not implemented")
}
