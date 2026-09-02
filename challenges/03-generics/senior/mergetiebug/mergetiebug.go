// Package mergetiebug — Gopher Workplace challenge.
package mergetiebug

import (
	"cmp"
)

// MergeSorted merges two key-sorted slices into one sorted slice.
// On a tie the element from a comes first.
//
// Examples:
//
//	MergeSorted(left, right, key) => one sorted slice
func MergeSorted[T any, K cmp.Ordered](a, b []T, key func(T) K) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if key(b[j]) <= key(a[i]) {
			out = append(out, b[j])
			j++
		} else {
			out = append(out, a[i])
			i++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
	// CHANGE CODE ABOVE THIS LINE
}
