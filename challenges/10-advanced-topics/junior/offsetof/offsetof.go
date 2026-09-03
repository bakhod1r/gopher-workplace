// Package offsetof — Gopher Workplace challenge.
package offsetof

import "unsafe"

// Rec is a record with mixed field widths.
type Rec struct {
	A byte
	B int64
	C byte
}

// Offsets returns the byte offset of each field of Rec from the start of
// the struct.
//
// Offsets are decided by the compiler from the field order and the
// alignment rules, not by the field sizes alone.
//
// Examples:
//
//	Offsets() => 0, 8, 16 for the declared layout
func Offsets() (a, b, c uintptr) {
	panic("not implemented")
}
