// Package extract pulls one byte out of a 64-bit word. Planted wrong shift stride.
package extract

// ByteAt returns byte n (0 = least significant) of v.
func ByteAt(v uint64, n uint) uint8 {
	// CHANGE CODE BELOW THIS LINE
	return uint8(v >> (4 * n))
	// CHANGE CODE ABOVE THIS LINE
}
