// Package maxof — Gopher Workplace challenge.
package maxof

import (
	"cmp"
)

// Max returns the larger of a and b. It returns b when they are equal.
//
// Examples:
//
//	Max(2, 5)          => 5
//	Max("a", "b")  => "b"
func Max[T cmp.Ordered](a, b T) T {
	// TODO(candidate): return the larger argument.
	panic("not implemented")
}
