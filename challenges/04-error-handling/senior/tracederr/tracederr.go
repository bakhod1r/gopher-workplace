// Package tracederr — Gopher Workplace challenge.
package tracederr

import (
	"errors"
	"fmt"
)

// ErrDisk is a stand-in low-level failure used by the tests.
var ErrDisk = errors.New("disk offline")

// TracedError carries an operation name alongside its cause.
type TracedError struct {
	Op    string
	Cause error
}

// Error returns the cause's message only.
func (e *TracedError) Error() string {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Unwrap exposes the cause.
func (e *TracedError) Unwrap() error {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Trace returns "<Op> -> <cause message>".
func (e *TracedError) Trace() string {
	// TODO(candidate): implement this.
	_ = fmt.Sprintf
	panic("not implemented")
}
