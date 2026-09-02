// Package topnstablebug — Gopher Workplace challenge.
package topnstablebug

import (
	"slices"
)

// TopN returns the n highest-scoring elements, best first.
// The input slice is left untouched and the result does not alias it.
//
// Examples:
//
//	TopN(rows, score, 2) => the two best rows
func TopN[T any](s []T, score func(T) int, n int) []T {
	// CHANGE CODE BELOW THIS LINE
	slices.SortStableFunc(s, func(a, b T) int {
		return score(b) - score(a)
	})
	if n > len(s) {
		n = len(s)
	}
	if n < 0 {
		n = 0
	}
	return s[:n]
	// CHANGE CODE ABOVE THIS LINE
}
