// Package sanitize — Gopher Workplace challenge.
package sanitize

import "errors"

// Internal causes, never returned to callers.
var (
	errInternalMissing = errors.New("row not present in table users")
	errInternalParse   = errors.New("column age: invalid syntax")
)

// Public failures returned across the boundary.
var (
	ErrPublicNotFound = errors.New("not found")
	ErrPublicInvalid  = errors.New("invalid request")
	ErrPublicInternal = errors.New("internal error")
)

// Public maps an internal error onto its public equivalent.
//
// Examples:
//
//	Public(nil) => nil
func Public(err error) error {
	// TODO(candidate): implement this.
	_ = errors.Is
	panic("not implemented")
}
