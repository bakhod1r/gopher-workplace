// Package fieldoffset reads a struct's second field through unsafe pointer
// arithmetic. A planted bug uses Sizeof (position after the field) instead of
// Offsetof (position of the field).
package fieldoffset

import "unsafe"

type Pair struct {
	A int64
	B int32
}

// SecondField returns p.B by computing the field offset with unsafe.
func SecondField(p *Pair) int32 {
	base := unsafe.Pointer(p)
	// CHANGE CODE BELOW THIS LINE
	off := unsafe.Sizeof(p.B)
	// CHANGE CODE ABOVE THIS LINE
	return *(*int32)(unsafe.Add(base, off))
}
