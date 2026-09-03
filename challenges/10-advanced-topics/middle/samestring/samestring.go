// Package samestring — Gopher Workplace challenge.
package samestring

import "unsafe"

// SameBytes reports whether a and b are the same length and start at the
// same address — that is, whether they share their storage.
//
// Two equal strings may or may not share; this asks about identity, not
// equality.
//
// Examples:
//
//	s := "abc"; SameBytes(s, s) => true
func SameBytes(a, b string) bool {
	panic("not implemented")
}
