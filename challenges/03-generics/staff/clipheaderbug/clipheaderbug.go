// Package clipheaderbug — Gopher Workplace challenge.
package clipheaderbug

import (
	"slices"
)

// Shrink returns the first n elements of s as a slice whose capacity
// equals its length.
//
// Appending to the result must never overwrite an element of s.
// n is clamped into [0, len(s)].
//
// Examples:
//
//	Shrink([]int{1, 2, 3, 4}, 2) => []int{1, 2} with cap 2
func Shrink[T any](s []T, n int) []T {
	// CHANGE CODE BELOW THIS LINE
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	return slices.Clip(s)[:n]
	// CHANGE CODE ABOVE THIS LINE
}
