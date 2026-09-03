// Package uint32at — Gopher Workplace challenge.
package uint32at

import "unsafe"

// Uint32At reads the native-endian uint32 at byte offset off in b.
//
// The read is bounds-checked and alignment-checked; anything out of range
// or misaligned reports false rather than faulting.
//
// Examples:
//
//	Uint32At(buf, 0) => the first four bytes as a uint32
func Uint32At(b []byte, off int) (uint32, bool) {
	panic("not implemented")
}
