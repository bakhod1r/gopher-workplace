// Package floatbits returns the raw IEEE-754 bits of a float64. A planted bug
// converts the VALUE (truncating to an integer) instead of reinterpreting the
// bit pattern via unsafe.Pointer.
package floatbits

import "unsafe"

// Bits returns the 64-bit pattern of f (same as math.Float64bits).
func Bits(f float64) uint64 {
	// CHANGE CODE BELOW THIS LINE
	g := float32(f)
	return uint64(*(*uint32)(unsafe.Pointer(&g)))
	// CHANGE CODE ABOVE THIS LINE
}
