// Package causeerr — Gopher Workplace challenge.
package causeerr

import (
	"errors"
	"fmt"
)

// ErrDB is a stand-in low-level failure used by the tests.
var ErrDB = errors.New("db unavailable")

// CodeError attaches a numeric code to an underlying failure.
type CodeError struct {
	Code  int
	Cause error
}

// Error implements the error interface as "[<Code>] <Cause>".
func (e *CodeError) Error() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}

// Unwrap exposes the underlying cause.
func (e *CodeError) Unwrap() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
