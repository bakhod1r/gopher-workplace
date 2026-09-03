// Package stringtobytes — Gopher Workplace challenge.
package stringtobytes

import "unsafe"

// Bytes returns a read-only byte view of s.
//
// The bytes belong to the string and may live in read-only memory, so the
// result must never be written to.
//
// Examples:
//
//	Bytes("hi") => []byte("hi"), sharing the string's bytes
func Bytes(s string) []byte {
	panic("not implemented")
}
