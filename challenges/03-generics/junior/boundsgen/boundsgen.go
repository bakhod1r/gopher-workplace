// Package boundsgen — Gopher Workplace challenge.
package boundsgen

import (
	"cmp"
)

// Bounds returns the smallest and largest elements of s and true.
// It returns zero values and false for an empty slice.
//
// Examples:
//
//	Bounds([]int{3, 1, 2}) => 1, 3, true
func Bounds[T cmp.Ordered](s []T) (T, T, bool) {
	// TODO(candidate): track the smallest and largest in one pass.
	panic("not implemented")
}
