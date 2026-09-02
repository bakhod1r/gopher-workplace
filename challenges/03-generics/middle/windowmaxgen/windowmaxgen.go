// Package windowmaxgen — Gopher Workplace challenge.
package windowmaxgen

import (
	"cmp"
)

// WindowMax returns the maximum of every window of n elements.
// It returns an empty result when n <= 0 or n > len(s).
func WindowMax[T cmp.Ordered](s []T, n int) []T {
	// TODO(candidate): track candidate indexes in a monotonic deque.
	panic("not implemented")
}
