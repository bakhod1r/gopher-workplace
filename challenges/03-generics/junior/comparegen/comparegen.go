// Package comparegen — Gopher Workplace challenge.
package comparegen

import (
	"cmp"
)

// Compare returns -1 when a < b, 0 when they are equal, and +1 when a > b.
//
// Examples:
//
//	Compare(1, 2) => -1
//	Compare(2, 2) => 0
//	Compare(3, 2) => 1
func Compare[T cmp.Ordered](a, b T) int {
	// TODO(candidate): return -1, 0, or +1.
	panic("not implemented")
}
