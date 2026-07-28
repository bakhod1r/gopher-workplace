// Package bankround rounds half to even ("banker's rounding").
// A planted bug rounds half away from zero instead.
package bankround

import "math"

// Round returns x rounded to the nearest integer, with ties going to the even
// neighbor (2.5 -> 2, 3.5 -> 4). Returned as float64.
func Round(x float64) float64 {
	// CHANGE CODE BELOW THIS LINE
	return math.Round(x)
	// CHANGE CODE ABOVE THIS LINE
}
