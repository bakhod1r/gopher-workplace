// Package absval returns absolute value for int8. Planted naive negate mishandles min.
package absval

// Abs returns the absolute value of x as an int (widened) so that the most
// negative int8 has a correct positive result.
func Abs(x int8) int {
	// CHANGE CODE BELOW THIS LINE
	if x < 0 {
		return int(-x)
	}
	return int(x)
	// CHANGE CODE ABOVE THIS LINE
}
