// Package floatbits extracts the unbiased exponent of a float64.
// A planted bug uses the wrong bias.
package floatbits

import "math"

// Exponent returns the unbiased base-2 exponent of a normal, positive x
// (x = mantissa * 2^exponent, 1 <= mantissa < 2). E.g. Exponent(1)=0,
// Exponent(2)=1, Exponent(0.5)=-1.
func Exponent(x float64) int {
	bits := math.Float64bits(x)
	raw := int((bits >> 52) & 0x7FF)
	// CHANGE CODE BELOW THIS LINE
	return raw - 1024
	// CHANGE CODE ABOVE THIS LINE
}
