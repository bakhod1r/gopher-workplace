// Package nanminmaxbug — Gopher Workplace challenge.
package nanminmaxbug

import (
	"cmp"
)

// MinMaxSkipNaN returns the smallest and largest element of xs.
// NaN entries are skipped. The bool reports whether any usable element was found.
//
// Examples:
//
//	MinMaxSkipNaN([]float64{math.NaN(), 3, 1}) => 1, 3, true
func MinMaxSkipNaN[T cmp.Ordered](xs []T) (T, T, bool) {
	// CHANGE CODE BELOW THIS LINE
	var mn, mx T
	seen := false
	for _, v := range xs {
		if !seen {
			mn, mx, seen = v, v, true
			continue
		}
		if v < mn {
			mn = v
		}
		if v > mx {
			mx = v
		}
	}
	return mn, mx, seen
	// CHANGE CODE ABOVE THIS LINE
}
