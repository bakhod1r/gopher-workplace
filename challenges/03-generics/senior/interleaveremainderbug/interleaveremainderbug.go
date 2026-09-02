// Package interleaveremainderbug — Gopher Workplace challenge.
package interleaveremainderbug

// Interleave alternates elements of a and b, then appends the longer tail.
//
// Examples:
//
//	Interleave([]int{1, 2, 3}, []int{9}) => []int{1, 9, 2, 3}
func Interleave[T any](a, b []T) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(a)+len(b))
	n := min(len(a), len(b))
	for i := 0; i < n; i++ {
		out = append(out, a[i], b[i])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
