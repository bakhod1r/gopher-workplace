// Package isany — Gopher Workplace challenge.
package isany

import "errors"

// Transient and fatal failures used by the tests.
var (
	ErrTimeout = errors.New("timeout")
	ErrReset   = errors.New("connection reset")
	ErrFatal   = errors.New("fatal")
)

// IsAny reports whether err matches any of the targets.
//
// Examples:
//
//	IsAny(ErrTimeout, ErrTimeout, ErrReset) => true
//	IsAny(ErrFatal, ErrTimeout)             => false
func IsAny(err error, targets ...error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
