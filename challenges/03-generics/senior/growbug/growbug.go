// Package growbug — Gopher Workplace challenge.
package growbug

import (
	"slices"
)

// Collect appends every value to s, reserving room first.
//
// Examples:
//
//	Collect([]int{1}, 2, 3) => []int{1, 2, 3}
func Collect[T any](s []T, vs ...T) []T {
	// CHANGE CODE BELOW THIS LINE
	grown := slices.Grow(s, len(vs))
	grown = append(grown, vs...)
	return s
	// CHANGE CODE ABOVE THIS LINE
}
