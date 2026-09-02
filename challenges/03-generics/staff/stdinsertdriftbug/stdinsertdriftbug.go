// Package stdinsertdriftbug — Gopher Workplace challenge.
package stdinsertdriftbug

import (
	"slices"
)

// InsertMarks inserts mark into s at each position in at.
// Positions are indices into the ORIGINAL slice and must be ascending.
// Out-of-range positions are skipped. The input is not modified.
//
// Examples:
//
//	InsertMarks([]int{1, 2, 3, 4}, []int{1, 3}, 0) => []int{1, 0, 2, 3, 0, 4}
func InsertMarks[T any](s []T, at []int, mark T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := slices.Clone(s)
	done := 0
	for _, p := range at {
		if p < 0 || p > len(s) {
			continue
		}
		out = slices.Insert(out, p, mark)
		done++
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
