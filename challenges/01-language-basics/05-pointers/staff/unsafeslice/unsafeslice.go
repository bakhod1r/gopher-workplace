// Package unsafeslice builds a []int32 view over an array via unsafe.Slice. A
// planted bug passes the length in BYTES instead of elements, creating an
// over-long slice.
package unsafeslice

import "unsafe"

// View returns a slice viewing all 4 elements of the array p points to.
func View(p *[4]int32) []int32 {
	// CHANGE CODE BELOW THIS LINE
	return unsafe.Slice(&p[0], unsafe.Sizeof(*p))
	// CHANGE CODE ABOVE THIS LINE
}
