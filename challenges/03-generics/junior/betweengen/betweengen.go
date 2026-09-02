// Package betweengen — Gopher Workplace challenge.
package betweengen

import (
	"cmp"
)

// Between reports whether lo <= v <= hi.
//
// Examples:
//
//	Between(2, 1, 3) => true
//	Between(1, 1, 3) => true
//	Between(4, 1, 3) => false
func Between[T cmp.Ordered](v, lo, hi T) bool {
	// TODO(candidate): report whether v lies within the inclusive range.
	panic("not implemented")
}
