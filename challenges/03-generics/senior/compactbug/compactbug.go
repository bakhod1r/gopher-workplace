// Package compactbug — Gopher Workplace challenge.
package compactbug

import (
	"cmp"
	"slices"
)

// Distinct returns the distinct elements of s in ascending order.
// The input is not modified.
//
// Examples:
//
//	Distinct([]int{3, 1, 3}) => []int{1, 3}
func Distinct[T cmp.Ordered](s []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := slices.Clone(s)
	if out == nil {
		out = []T{}
	}
	return slices.Compact(out)
	// CHANGE CODE ABOVE THIS LINE
}
