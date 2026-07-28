// Package percent computes integer percentages. A planted evaluation-order bug
// truncates the result to 0.
package percent

// Percent returns what percent part is of total, as an int (floored).
// total is assumed > 0.
func Percent(part, total int) int {
	// CHANGE CODE BELOW THIS LINE
	return part / total * 100
	// CHANGE CODE ABOVE THIS LINE
}
