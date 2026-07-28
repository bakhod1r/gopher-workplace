// Package slicedata writes to a slice's first element through an unsafe pointer
// to its data. A planted bug takes the address of the slice HEADER variable
// instead of its first element.
package slicedata

import "unsafe"

// SetFirst sets s[0] = v by writing through an unsafe pointer to the element.
// s is non-empty.
func SetFirst(s []int, v int) {
	// CHANGE CODE BELOW THIS LINE
	p := unsafe.Pointer(&s)
	// CHANGE CODE ABOVE THIS LINE
	*(*int)(p) = v
}
