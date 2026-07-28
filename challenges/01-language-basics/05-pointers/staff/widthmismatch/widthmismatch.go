// Package widthmismatch reads a uint64 from an 8-byte value via unsafe. A planted
// bug reinterprets through *uint32, reading only the low 4 bytes.
package widthmismatch

import "unsafe"

// AsU64 returns the raw 64 bits of x reinterpreted as uint64.
func AsU64(x int64) uint64 {
	// CHANGE CODE BELOW THIS LINE
	return uint64(*(*uint32)(unsafe.Pointer(&x)))
	// CHANGE CODE ABOVE THIS LINE
}
