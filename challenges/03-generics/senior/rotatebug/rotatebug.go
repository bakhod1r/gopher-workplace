// Package rotatebug — Gopher Workplace challenge.
package rotatebug

// Rotate returns s rotated left by k. Negative k rotates right.
// The input is not modified.
//
// Examples:
//
//	Rotate([]int{1, 2, 3}, -1) => []int{3, 1, 2}
func Rotate[T any](s []T, k int) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(s))
	if len(s) == 0 {
		return out
	}
	k = k % len(s)
	out = append(out, s[k:]...)
	out = append(out, s[:k]...)
	return out
	// CHANGE CODE ABOVE THIS LINE
}
