// Package inrange — Gopher Workplace challenge.
package inrange

import "errors"

// Range failures.
var (
	ErrBadBounds  = errors.New("lo greater than hi")
	ErrOutOfRange = errors.New("value out of range")
)

// InRange reports whether n lies within the inclusive range [lo, hi].
//
// Examples:
//
//	InRange(5, 1, 10)  => nil
//	InRange(5, 10, 1)  => ErrBadBounds
func InRange(n, lo, hi int) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
