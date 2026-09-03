// Package slicelen — Gopher Workplace challenge.
package slicelen

import "unsafe"

// Words returns a []uint32 view over b's bytes.
//
// unsafe.Slice takes a count of elements, not of bytes: passing the byte
// length produces a view four times too long, running off the end of the
// buffer.
//
// Examples:
//
//	Words(eightBytes) => a 2-element view, true
func Words(b []byte) ([]uint32, bool) {
	// CHANGE CODE BELOW THIS LINE
	if len(b) == 0 || len(b)%4 != 0 {
		return nil, false
	}
	p := unsafe.Pointer(unsafe.SliceData(b))
	if uintptr(p)&3 != 0 {
		return nil, false
	}
	return unsafe.Slice((*uint32)(p), len(b)), true
	// CHANGE CODE ABOVE THIS LINE
}
