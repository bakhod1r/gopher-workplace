// Package firstof — Gopher Workplace challenge.
package firstof

// First returns the first element of s and true.
// It returns the zero value and false when s is empty.
//
// Examples:
//
//	First([]int{3, 1})  => 3, true
//	First([]int{})      => 0, false
func First[T any](s []T) (T, bool) {
	// TODO(candidate): guard the empty slice, then return s[0].
	panic("not implemented")
}
