// Package lowerboundbug — Gopher Workplace challenge.
package lowerboundbug

import (
	"cmp"
)

// LowerBound returns the index of the first element of the sorted
// slice s that is greater than or equal to v.
//
// It returns len(s) when no such element exists, and must run in
// logarithmic time.
//
// Examples:
//
//	LowerBound([]int{1, 2, 2, 2, 3}, 2) => 1
func LowerBound[T cmp.Ordered](s []T, v T) int {
	// CHANGE CODE BELOW THIS LINE
	lo, hi := 0, len(s)
	for lo < hi {
		mid := int(uint(lo+hi) >> 1)
		if s[mid] <= v {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
	// CHANGE CODE ABOVE THIS LINE
}
