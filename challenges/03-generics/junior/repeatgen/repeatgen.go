// Package repeatgen — Gopher Workplace challenge.
package repeatgen

// Repeat returns s concatenated with itself n times.
// It returns an empty slice when n <= 0.
//
// Examples:
//
//	Repeat([]int{1, 2}, 2) => []int{1, 2, 1, 2}
//	Repeat([]int{1}, 0)    => []int{}
func Repeat[T any](s []T, n int) []T {
	// TODO(candidate): append a copy of s n times.
	panic("not implemented")
}
