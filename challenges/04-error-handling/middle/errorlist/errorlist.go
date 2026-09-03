// Package errorlist — Gopher Workplace challenge.
package errorlist

import (
	"errors"
	"strings"
)

// Rule failures used by the tests.
var (
	ErrA = errors.New("rule a")
	ErrB = errors.New("rule b")
)

// Errors is a collection of failures reported together.
type Errors []error

// Error joins the member messages with "; ".
func (e Errors) Error() string {
	// TODO(candidate): implement this.
	_ = strings.Join
	panic("not implemented")
}

// Unwrap exposes the members to errors.Is and errors.As.
func (e Errors) Unwrap() []error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
