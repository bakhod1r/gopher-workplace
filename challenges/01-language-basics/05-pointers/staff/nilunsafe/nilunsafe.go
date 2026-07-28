// Package nilunsafe reads an int through an unsafe pointer, defaulting on nil. A
// planted bug reads before the nil check, dereferencing a nil pointer.
package nilunsafe

import "unsafe"

// ReadOr returns *(*int)(p), or def when p is nil.
func ReadOr(p unsafe.Pointer, def int) int {
	// CHANGE CODE BELOW THIS LINE
	v := *(*int)(p)
	if p == nil {
		return def
	}
	return v
	// CHANGE CODE ABOVE THIS LINE
}
