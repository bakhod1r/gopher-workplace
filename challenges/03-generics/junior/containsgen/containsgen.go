// Package containsgen — Gopher Workplace challenge.
package containsgen

// Contains reports whether v appears in s.
//
// Examples:
//
//	Contains([]int{1, 2, 3}, 2) => true
//	Contains([]int{1, 2, 3}, 9) => false
func Contains[T comparable](s []T, v T) bool {
	// TODO(candidate): scan the slice for an element equal to v.
	panic("not implemented")
}
