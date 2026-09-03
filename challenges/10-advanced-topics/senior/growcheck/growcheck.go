// Package growcheck — Gopher Workplace challenge.
package growcheck

import "unsafe"

// Grew reports whether after occupies different storage from before —
// that is, whether the append that produced it had to reallocate.
//
// Examples:
//
//	s := make([]int, 0, 1); Grew(s, append(s, 1)) => false
func Grew(before, after []int) bool {
	panic("not implemented")
}
