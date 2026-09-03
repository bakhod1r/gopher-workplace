// Package classify — Gopher Workplace challenge.
package classify

import "errors"

// Domain failures used by the tests.
var (
	ErrNotFound = errors.New("not found")
	ErrDenied   = errors.New("permission denied")
	ErrConflict = errors.New("conflict")
)

// Status maps err onto an HTTP status code.
//
// Examples:
//
//	Status(nil)         => 200
//	Status(ErrNotFound) => 404
func Status(err error) int {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
