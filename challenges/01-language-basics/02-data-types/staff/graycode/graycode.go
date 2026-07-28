// Package graycode converts a binary number to reflected Gray code.
// A planted bug shifts the wrong direction.
package graycode

// ToGray returns the Gray-code encoding of x: x XOR (x >> 1).
func ToGray(x uint32) uint32 {
	// CHANGE CODE BELOW THIS LINE
	return x ^ (x << 1)
	// CHANGE CODE ABOVE THIS LINE
}
