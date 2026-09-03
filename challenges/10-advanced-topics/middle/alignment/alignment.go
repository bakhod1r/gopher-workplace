// Package alignment — Gopher Workplace challenge.
package alignment

import "unsafe"

// Aligned reports whether b's first byte sits at an address that is a
// multiple of n.
//
// n must be a power of two; anything else, or an empty slice, reports
// false.
//
// Examples:
//
//	Aligned(buf, 8) => true when buf starts on an 8-byte boundary
func Aligned(b []byte, n uintptr) bool {
	panic("not implemented")
}
