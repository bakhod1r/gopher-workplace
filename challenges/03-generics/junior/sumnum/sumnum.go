// Package sumnum — Gopher Workplace challenge.
package sumnum

// Number is the set of numeric types this package works with.
type Number interface {
	~int | ~int64 | ~float64
}

// Sum returns the total of s. It returns 0 for an empty slice.
//
// Examples:
//
//	Sum([]int{1, 2, 3})       => 6
//	Sum([]float64{0.5, 0.5}) => 1
func Sum[T Number](s []T) T {
	// TODO(candidate): add every element into a running total.
	panic("not implemented")
}
