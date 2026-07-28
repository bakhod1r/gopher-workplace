// Package money converts dollars to integer cents. Planted truncation order bug.
package money

// Cents returns the number of whole cents in a dollar amount, rounding down.
func Cents(dollars float64) int64 {
	// CHANGE CODE BELOW THIS LINE
	return int64(dollars) * 100
	// CHANGE CODE ABOVE THIS LINE
}
