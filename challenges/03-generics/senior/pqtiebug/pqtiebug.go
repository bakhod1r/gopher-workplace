// Package pqtiebug — Gopher Workplace challenge.
package pqtiebug

import (
	"cmp"
)

// InsertSorted places v into s, which is already sorted by key.
// Equal keys keep arrival order: v goes after existing equals.
//
// Examples:
//
//	InsertSorted(jobs, j, prio) => j after every job of the same priority
func InsertSorted[T any, K cmp.Ordered](s []T, v T, key func(T) K) []T {
	// CHANGE CODE BELOW THIS LINE
	k := key(v)
	i := 0
	for i < len(s) && key(s[i]) < k {
		i++
	}
	s = append(s, v)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
	// CHANGE CODE ABOVE THIS LINE
}
