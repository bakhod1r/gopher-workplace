// Package windowmaxexpirebug — Gopher Workplace challenge.
package windowmaxexpirebug

import (
	"cmp"
)

// WindowMax returns the maximum of each window of width w.
// It returns an empty slice when w is out of range.
//
// Examples:
//
//	WindowMax([]int{1, 3, 2}, 2) => []int{3, 3}
func WindowMax[T cmp.Ordered](s []T, w int) []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0)
	if w <= 0 || w > len(s) {
		return out
	}
	dq := make([]int, 0, len(s))
	for i := range s {
		for len(dq) > 0 && s[dq[len(dq)-1]] <= s[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if i >= w-1 {
			out = append(out, s[dq[0]])
		}
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
