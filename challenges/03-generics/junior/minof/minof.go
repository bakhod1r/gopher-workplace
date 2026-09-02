// Package minof — Gopher Workplace challenge.
package minof

import (
	"cmp"
)

// Min returns the smaller of a and b. It returns b when they are equal.
//
// Examples:
//
//	Min(2, 5) => 2
//	Min(3.5, 1.5) => 1.5
func Min[T cmp.Ordered](a, b T) T {
	// TODO(candidate): return the smaller argument.
	panic("not implemented")
}
