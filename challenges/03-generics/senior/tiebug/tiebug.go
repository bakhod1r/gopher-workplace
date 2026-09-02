// Package tiebug — Gopher Workplace challenge.
package tiebug

import (
	"cmp"
)

// MaxBy returns the element with the largest key and true.
// On a tie the earlier element wins.
//
// Examples:
//
//	MaxBy(rows, score) => the first best-scoring row
func MaxBy[T any, K cmp.Ordered](s []T, key func(T) K) (T, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	best, bestKey := s[0], key(s[0])
	for _, v := range s[1:] {
		k := key(v)
		if k >= bestKey {
			best, bestKey = v, k
		}
	}
	return best, true
	// CHANGE CODE ABOVE THIS LINE
}
