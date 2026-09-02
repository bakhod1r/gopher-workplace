// Package mediangen — Gopher Workplace challenge.
package mediangen

// Float is the set of floating-point types.
type Float interface {
	~float32 | ~float64
}

// Median returns the middle value of s and true.
// For an even count it averages the two middle values.
// It returns the zero value and false for an empty slice.
func Median[T Float](s []T) (T, bool) {
	// TODO(candidate): sort a copy, then read the middle.
	panic("not implemented")
}
