// Package structlayout — Gopher Workplace challenge.
package structlayout

import "unsafe"

// Sizer reports its own size in bytes.
type Sizer interface {
	Size() uintptr
}

// Padded interleaves small and large fields, so the compiler inserts padding.
type Padded struct {
	A bool  // 1 byte, then 7 bytes of padding
	B int64 // 8 bytes
	C bool  // 1 byte, then 7 bytes of padding
}

// Size returns the size of the struct.
func (p Padded) Size() uintptr {
	// TODO(candidate): unsafe.Sizeof of the receiver's type.
	panic("not implemented")
}

// Packed holds the same fields, ordered to minimise padding.
type Packed struct {
	// TODO(candidate): the same fields (an int64 and two bools), reordered
	// so the struct is smaller than Padded.
	B int64
	A bool
	C bool
}

// Size returns the size of the struct.
func (p Packed) Size() uintptr {
	// TODO(candidate): unsafe.Sizeof of the receiver's type.
	panic("not implemented")
}

// TotalBytes reports the bytes a slice of n such records occupies.
//
// Examples:
//
//	TotalBytes(Packed{}, 1000) => 1000 * unsafe.Sizeof(Packed{})
func TotalBytes(s Sizer, n int) uintptr {
	// TODO(candidate): size times count.
	panic("not implemented")
}

var _ = unsafe.Sizeof(0)
