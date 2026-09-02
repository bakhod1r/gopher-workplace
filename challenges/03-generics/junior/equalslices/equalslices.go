// Package equalslices — Gopher Workplace challenge.
package equalslices

// Equal reports whether a and b have the same length and
// the same elements in the same order.
//
// Examples:
//
//	Equal([]int{1, 2}, []int{1, 2}) => true
//	Equal([]int{1, 2}, []int{2, 1}) => false
func Equal[T comparable](a, b []T) bool {
	// TODO(candidate): compare lengths, then element by element.
	panic("not implemented")
}
