// Package noallocerr — Gopher Workplace challenge.
package noallocerr

import "errors"

// ErrEmpty reports empty input.
var ErrEmpty = errors.New("empty input")

// Check validates s without allocating.
//
// Examples:
//
//	Check("")  => ErrEmpty
//	Check("a") => nil
func Check(s string) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
