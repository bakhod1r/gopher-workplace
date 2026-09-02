// Package lastof — Gopher Workplace challenge.
package lastof

// Last returns the final element of s and true.
// It returns the zero value and false when s is empty.
//
// Examples:
//
//	Last([]int{3, 1, 4}) => 4, true
//	Last([]int{})        => 0, false
func Last[T any](s []T) (T, bool) {
	// TODO(candidate): guard the empty slice, then index the last position.
	panic("not implemented")
}
