// Package endianswap reverses the byte order of a uint32.
// A planted bug uses the wrong mask on one byte lane.
package endianswap

// Swap32 returns x with its four bytes reversed (big-endian <-> little-endian).
func Swap32(x uint32) uint32 {
	// CHANGE CODE BELOW THIS LINE
	return x<<24 | (x<<8)&0xFF00 | (x>>8)&0xFF00 | x>>24
	// CHANGE CODE ABOVE THIS LINE
}
