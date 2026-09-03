// Package customis — Gopher Workplace challenge.
package customis

import "fmt"

// StatusError reports a transport status code.
type StatusError struct {
	Code int
}

// Error implements the error interface as "status <Code>".
func (e *StatusError) Error() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}

// Is reports whether target is the same code or the class marker for it.
func (e *StatusError) Is(target error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
