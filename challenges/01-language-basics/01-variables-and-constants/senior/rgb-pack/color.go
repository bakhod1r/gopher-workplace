// Package color packs 8-bit RGB into a uint32. A planted shift bug swaps a channel.
package color

// Pack encodes r,g,b (each 0..255) into 0x00RRGGBB.
func Pack(r, g, b uint8) uint32 {
	// CHANGE CODE BELOW THIS LINE
	return uint32(r)<<8 | uint32(g)<<8 | uint32(b)
	// CHANGE CODE ABOVE THIS LINE
}

// Red extracts the red channel from a packed value.
func Red(v uint32) uint8 { return uint8(v >> 16) }
