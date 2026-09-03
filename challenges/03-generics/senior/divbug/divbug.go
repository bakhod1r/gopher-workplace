// Package divbug — Gopher Workplace challenge.
package divbug

// Integer is the set of signed integer types used here.
type Integer interface {
	~int | ~int64
}

// SumRange returns lo + (lo+1) + ... + hi.
// It returns 0 when hi < lo.
//
// Examples:
//
//	SumRange(1, 4) => 10
func SumRange[T Integer](lo, hi T) T {
	// CHANGE CODE BELOW THIS LINE
	if hi < lo {
		var zero T
		return zero
	}
	n := hi - lo + 1
	return (lo + hi) / 2 * n
	// CHANGE CODE ABOVE THIS LINE
}
