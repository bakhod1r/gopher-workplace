// Package seedbug — Gopher Workplace challenge.
package seedbug

import (
	"cmp"
)

// MinOf returns the smallest element and true.
// It returns the zero value and false for an empty slice.
//
// Examples:
//
//	MinOf([]int{4, 7}) => 4, true
func MinOf[T cmp.Ordered](s []T) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	var best T
	for _, v := range s {
		if v < best {
			best = v
		}
	}
	return best, true
	// CHANGE CODE ABOVE THIS LINE
}
