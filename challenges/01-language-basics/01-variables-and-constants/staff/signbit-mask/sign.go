// Package sign inspects the sign bit of a byte. Planted wrong mask.
package sign

// Negative reports whether the int8 value stored in b (as its raw bits) is
// negative — i.e. whether the sign bit (bit 7) is set.
func Negative(b uint8) bool {
	// CHANGE CODE BELOW THIS LINE
	return b&0x40 != 0
	// CHANGE CODE ABOVE THIS LINE
}
