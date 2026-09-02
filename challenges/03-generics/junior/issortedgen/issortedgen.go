// Package issortedgen — Gopher Workplace challenge.
package issortedgen

import (
	"cmp"
)

// IsSorted reports whether s is in non-decreasing order.
// Slices of length 0 or 1 are sorted.
//
// Examples:
//
//	IsSorted([]int{1, 2, 2}) => true
//	IsSorted([]int{2, 1})    => false
func IsSorted[T cmp.Ordered](s []T) bool {
	// TODO(candidate): check every adjacent pair.
	panic("not implemented")
}
