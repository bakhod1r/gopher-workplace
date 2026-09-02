// Package scaleroundbug — Gopher Workplace challenge.
package scaleroundbug

// Integer is the set of integer types this helper accepts.
type Integer interface {
	~int | ~int32 | ~int64
}

// ScaleAll returns each value scaled by pct percent, rounded half away from zero.
//
// Examples:
//
//	ScaleAll([]int{250}, 15) => []int{38}
func ScaleAll[T Integer](vals []T, pct T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, len(vals))
	for i, v := range vals {
		out[i] = (v / 100) * pct
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
