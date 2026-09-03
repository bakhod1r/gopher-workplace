// Package structroundtrip — Gopher Workplace challenge.
package structroundtrip

import "unsafe"

// Frame is a fixed-layout record of scalars.
type Frame struct {
	Kind  uint32
	Seq   uint32
	Stamp int64
}

// Encode returns a byte view of f, for the tests to feed back in.
func Encode(f *Frame) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(f)), unsafe.Sizeof(*f))
}

// Decode reinterprets b as a Frame, copying it out so the result does not
// alias b.
//
// The length must be exactly the frame's size and the start must be
// correctly aligned; otherwise the second result is false.
//
// Examples:
//
//	Decode(encoded) => the frame, true
func Decode(b []byte) (Frame, bool) {
	panic("not implemented")
}
