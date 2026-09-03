// Package unwrapboth — Gopher Workplace challenge.
package unwrapboth

import "errors"

// Sample failures used by the tests.
var (
	ErrA = errors.New("a")
	ErrB = errors.New("b")
)

// Multi carries a primary failure alongside secondary ones.
type Multi struct {
	First  error
	Others []error
}

// Error returns the primary failure's message.
func (m *Multi) Error() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Unwrap exposes the primary and the others for matching.
func (m *Multi) Unwrap() []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Primary returns the main failure.
func (m *Multi) Primary() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
