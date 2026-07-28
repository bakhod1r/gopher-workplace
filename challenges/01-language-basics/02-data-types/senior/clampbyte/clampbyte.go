// Package clampbyte saturates an int to a byte (0..255) for pixel math.
// A planted bug wraps instead of clamping.
package clampbyte

// Clamp converts x to a byte, saturating out-of-range values to 0 or 255
// (as image processing requires), never wrapping.
func Clamp(x int) byte {
	// CHANGE CODE BELOW THIS LINE
	return byte(x)
	// CHANGE CODE ABOVE THIS LINE
}
