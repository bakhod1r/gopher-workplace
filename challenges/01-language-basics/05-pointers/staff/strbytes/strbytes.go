// Package strbytes reads a string's length from its header via unsafe, without
// calling len. A planted bug reads the DATA pointer word instead of the LENGTH
// word of the header.
package strbytes

import "unsafe"

type strHeader struct {
	Data unsafe.Pointer
	Len  int
}

// ByteLen returns len(s) by reading the string header fields directly.
func ByteLen(s string) int {
	h := (*strHeader)(unsafe.Pointer(&s))
	// CHANGE CODE BELOW THIS LINE
	return int(uintptr(h.Data))
	// CHANGE CODE ABOVE THIS LINE
}
