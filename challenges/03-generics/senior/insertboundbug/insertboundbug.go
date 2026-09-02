// Package insertboundbug — Gopher Workplace challenge.
package insertboundbug

import (
	"slices"
)

// InsertAt returns s with v inserted at index i.
// Valid indexes run 0..len(s); an out-of-range index returns s.
//
// Examples:
//
//	InsertAt([]int{1, 3}, 1, 2) => []int{1, 2, 3}
func InsertAt[T any](s []T, i int, v T) []T {
	// CHANGE CODE BELOW THIS LINE
	if i < 0 || i >= len(s) {
		return s
	}
	return slices.Insert(slices.Clone(s), i, v)
	// CHANGE CODE ABOVE THIS LINE
}
