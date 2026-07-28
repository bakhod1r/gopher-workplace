// Package ptrfieldsize returns the size of a struct that holds two pointers. A
// planted bug measures a single pointer instead of the whole struct.
package ptrfieldsize

import "unsafe"

type Pair struct {
	A *int
	B *int
}

// Size returns the byte size of Pair (two 8-byte pointers = 16 on 64-bit).
func Size() uintptr {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Sizeof((*int)(nil))
	// CHANGE CODE ABOVE THIS LINE
}
