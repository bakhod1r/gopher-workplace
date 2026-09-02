// Package sortedcopy — Gopher Workplace challenge.
package sortedcopy

import (
	"cmp"
)

// Sorted returns a sorted copy of s. The input is not modified.
//
// Examples:
//
//	Sorted([]int{3, 1, 2}) => []int{1, 2, 3}
func Sorted[T cmp.Ordered](s []T) []T {
	// TODO(candidate): copy s, then sort the copy.
	panic("not implemented")
}
