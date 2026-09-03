// Package clipbug — Gopher Workplace challenge.
package clipbug

// Head returns the first n elements of s as an independent view.
// Appending to the result must not touch s.
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
	return s[:n]
	// CHANGE CODE ABOVE THIS LINE
}
