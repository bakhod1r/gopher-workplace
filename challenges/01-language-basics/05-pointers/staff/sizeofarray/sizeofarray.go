// Package sizeofarray returns the total byte size of a [4]int32 array. A planted
// bug measures one element instead of the whole array.
package sizeofarray

import "unsafe"

// TotalSize returns the byte size of the entire array p points to.
func TotalSize(p *[4]int32) uintptr {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Sizeof(p[0])
	// CHANGE CODE ABOVE THIS LINE
}
