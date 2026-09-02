// Package lesserselfbug — Gopher Workplace challenge.
package lesserselfbug

// Lesser is satisfied by types that can order themselves against their own type.
type Lesser[T any] interface {
	Less(T) bool
}

// MinOf returns the smallest element of xs according to its own Less method.
// The bool reports whether xs held anything.
//
// Examples:
//
//	MinOf([]ver{{1, 5}, {1, 2}}) => ver{1, 2}, true
func MinOf[T Lesser[T]](xs []T) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(xs) == 0 {
		var zero T
		return zero, false
	}
	best := xs[0]
	for _, v := range xs[1:] {
		if best.Less(v) {
			best = v
		}
	}
	return best, true
	// CHANGE CODE ABOVE THIS LINE
}
