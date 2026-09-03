// Package binsearchbug — Gopher Workplace challenge.
package binsearchbug

import (
	"cmp"
)

// SearchBy returns the index of the first element whose key equals
// target, or the insertion point and false.
//
// Examples:
//
//	SearchBy(rows, idOf, 1) => index of the first id 1
func SearchBy[T any, K cmp.Ordered](s []T, key func(T) K, target K) (int, bool) {
	// CHANGE CODE BELOW THIS LINE
	lo, hi := 0, len(s)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if key(s[mid]) <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if lo < len(s) && key(s[lo]) == target {
		return lo, true
	}
	return lo, false
	// CHANGE CODE ABOVE THIS LINE
}
