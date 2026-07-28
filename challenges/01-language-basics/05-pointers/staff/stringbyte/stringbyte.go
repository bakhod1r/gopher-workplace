// Package stringbyte returns the first byte of a string via unsafe, without
// allocating a []byte copy. A planted bug takes the address of the string HEADER
// variable instead of its data.
package stringbyte

import "unsafe"

// FirstByte returns s[0] read through the string's data pointer. s is non-empty.
func FirstByte(s string) byte {
	// CHANGE CODE BELOW THIS LINE
	p := unsafe.Pointer(&s)
	// CHANGE CODE ABOVE THIS LINE
	return *(*byte)(p)
}
