// Package isnotfound — Gopher Workplace challenge.
package isnotfound

import "errors"

// Sentinel failures used by the tests.
var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("permission denied")
)

// IsNotFound reports whether err was ultimately caused by ErrNotFound.
//
// Examples:
//
//	IsNotFound(ErrNotFound) => true
//	IsNotFound(ErrDenied)   => false
func IsNotFound(err error) bool {
	// TODO(candidate): implement this.
	panic("not implemented")
}
