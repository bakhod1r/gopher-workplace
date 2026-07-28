// Package offsetpad reads the int64 field that follows a bool, through unsafe.
// A planted bug assumes the field sits at offset 1 (right after the bool),
// ignoring the 7 bytes of alignment padding.
package offsetpad

import "unsafe"

type Rec struct {
	Flag bool
	N    int64
}

// ReadN returns r.N read via an unsafe offset from the struct base.
func ReadN(r *Rec) int64 {
	base := unsafe.Pointer(r)
	// CHANGE CODE BELOW THIS LINE
	off := unsafe.Sizeof(r.Flag)
	// CHANGE CODE ABOVE THIS LINE
	return *(*int64)(unsafe.Add(base, off))
}
