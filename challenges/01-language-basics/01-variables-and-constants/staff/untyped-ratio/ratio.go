// Package ratio holds a compile-time constant ratio. Planted typed truncation.
package ratio

// GoldenApprox approximates the golden ratio as an untyped constant so it keeps
// full precision until used. A planted bug forces integer truncation.
// CHANGE CODE BELOW THIS LINE
const GoldenApprox = 233 / 144

// CHANGE CODE ABOVE THIS LINE

// Value returns the ratio as a float64.
func Value() float64 { return GoldenApprox }
