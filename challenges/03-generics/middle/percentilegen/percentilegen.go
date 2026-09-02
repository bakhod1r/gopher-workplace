// Package percentilegen — Gopher Workplace challenge.
package percentilegen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Percentile returns the nearest-rank percentile of s and true.
// p is clamped to [0, 100]. It returns zero and false when s is empty.
func Percentile[T Float](s []T, p float64) (T, bool) {
	// TODO(candidate): sort a copy, then pick the nearest rank.
	panic("not implemented")
}
