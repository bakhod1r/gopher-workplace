// Package opaqueerr — Gopher Workplace challenge.
package opaqueerr

import "errors"

// Internal causes; callers use the predicates instead.
var (
	ErrTransient = errors.New("transient failure")
	ErrInvalid   = errors.New("invalid input")
)

// IsRetryable reports whether the operation may be retried.
//
// Examples:
//
//	IsRetryable(ErrTransient) => true
func IsRetryable(err error) bool {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}

// IsRejected reports whether the input was rejected.
//
// Examples:
//
//	IsRejected(ErrInvalid) => true
func IsRejected(err error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
