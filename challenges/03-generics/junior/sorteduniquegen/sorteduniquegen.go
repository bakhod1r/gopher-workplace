// Package sorteduniquegen — Gopher Workplace challenge.
package sorteduniquegen

import (
	"cmp"
)

// SortedUnique returns the distinct elements of s in ascending order.
//
// Examples:
//
//	SortedUnique([]int{3, 1, 3}) => []int{1, 3}
func SortedUnique[T cmp.Ordered](s []T) []T {
	// TODO(candidate): drop duplicates, then sort.
	panic("not implemented")
}
