// Package percentilebug — Gopher Workplace challenge.
package percentilebug

import (
	"math"
	"slices"
)

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Percentile returns the nearest-rank percentile and true.
// p is clamped to [0, 100].
//
// Examples:
//
//	Percentile([]float64{1, 2, 3, 4}, 50) => 2, true
func Percentile[T Float](s []T, p float64) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	if p < 0 {
		p = 0
	}
	if p > 100 {
		p = 100
	}
	c := slices.Clone(s)
	slices.Sort(c)
	rank := int(math.Ceil(p / 100 * float64(len(c))))
	return c[rank-1], true
	// CHANGE CODE ABOVE THIS LINE
}
