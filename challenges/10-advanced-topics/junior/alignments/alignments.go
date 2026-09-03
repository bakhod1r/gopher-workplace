// Package alignments — Gopher Workplace challenge.
package alignments

import "unsafe"

// Alignments returns the alignment requirement of a byte, an int32, an
// int64 and a string.
//
// A type's alignment is the boundary its address must be a multiple of.
//
// Examples:
//
//	Alignments() => 1, 4, 8, 8 on a 64-bit build
func Alignments() (b, i32, i64, s uintptr) {
	panic("not implemented")
}
