// Package retainbug — Gopher Workplace challenge.
package retainbug

// Head returns the first n elements of s.
//
// The result must be independent of s: it may neither observe later
// writes to s nor keep s's backing array alive.
// n is clamped into [0, len(s)].
//
// Examples:
//
//	Head([]int{1, 2, 3}, 2) => []int{1, 2}
func Head[T any](s []T, n int) []T {
	// CHANGE CODE BELOW THIS LINE
	if n < 0 {
		n = 0
	}
	if n > len(s) {
		n = len(s)
	}
	return s[:n:n]
	// CHANGE CODE ABOVE THIS LINE
}
