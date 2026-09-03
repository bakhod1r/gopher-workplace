// Package isprecedence — Gopher Workplace challenge.
package isprecedence

import (
	"errors"
	"fmt"
)

// ErrBase is a stand-in root failure used by the tests.
var ErrBase = errors.New("base")

// CodedError attaches a code to a cause.
type CodedError struct {
	Code  int
	Cause error
}

// Error implements the error interface as "code <Code>: <Cause>".
func (e *CodedError) Error() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}

// Is matches another *CodedError carrying the same code.
func (e *CodedError) Is(target error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Unwrap exposes the cause.
func (e *CodedError) Unwrap() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
