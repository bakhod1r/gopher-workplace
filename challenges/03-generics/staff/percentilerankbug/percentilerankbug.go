// Package percentilerankbug — Gopher Workplace challenge.
package percentilerankbug

import (
	"cmp"
	"slices"
)

// Percentile returns the nearest-rank p-th percentile of xs, for p in [0, 100].
// The input slice is not modified. The bool reports whether the request was valid.
//
// Examples:
//
//	Percentile([]int{1, 2, 3, 4, 5}, 50) => 3, true
func Percentile[T cmp.Ordered](xs []T, p int) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	var zero T
	if len(xs) == 0 || p < 0 || p > 100 {
		return zero, false
	}
	sorted := slices.Clone(xs)
	slices.Sort(sorted)
	idx := (p*len(sorted) + 99) / 100
	if idx < 0 {
		idx = 0
	}
	if idx >= len(sorted) {
		idx = len(sorted) - 1
	}
	return sorted[idx], true
	// CHANGE CODE ABOVE THIS LINE
}
