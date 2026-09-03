// Package nanbug — Gopher Workplace challenge.
package nanbug

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// MinIgnoringNaN returns the smallest non-NaN element and true.
// It returns zero and false when there is no such element.
//
// Examples:
//
//	MinIgnoringNaN([]float64{NaN, 2}) => 2, true
func MinIgnoringNaN[T Float](s []T) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best := s[0]
	for _, v := range s[1:] {
		if v != v {
			continue
		}
		if v < best {
			best = v
		}
	}
	return best, true
	// CHANGE CODE ABOVE THIS LINE
}
