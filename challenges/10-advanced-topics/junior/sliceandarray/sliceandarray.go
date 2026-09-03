// Package sliceandarray — Gopher Workplace challenge.
package sliceandarray

import "unsafe"

// Sizes returns the size of a [8]int array and of an []int slice header.
//
// The array's size is its contents; the slice's is three words, whatever it
// points at.
//
// Examples:
//
//	Sizes() => 64, 24 on a 64-bit build
func Sizes() (arr, sl uintptr) {
	panic("not implemented")
}
