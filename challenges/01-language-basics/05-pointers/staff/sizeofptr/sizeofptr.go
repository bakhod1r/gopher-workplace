// Package sizeofptr returns the size of the pointed-to array's element type. A
// planted bug measures the POINTER (8 bytes) instead of the element it points to.
package sizeofptr

import "unsafe"

// ElemSize returns the size in bytes of a single element of the array p points
// to.
func ElemSize(p *[8]int32) uintptr {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Sizeof(p)
	// CHANGE CODE ABOVE THIS LINE
}
