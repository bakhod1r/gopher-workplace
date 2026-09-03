// Package samearray — Gopher Workplace challenge.
package samearray

import "unsafe"

// SameArray reports whether a and b start at the same element of the
// same backing array.
//
// Comparing slices with == is not allowed; comparing their data pointers
// is.
//
// Examples:
//
//	s := []int{1, 2}; SameArray(s, s[:1]) => true
func SameArray(a, b []int) bool {
	panic("not implemented")
}
