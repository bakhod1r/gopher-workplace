// Package checkedadd adds int64 values with overflow detection.
// A planted bug misses positive overflow.
package checkedadd

import "math"

// Add returns a+b and ok=false if the mathematical sum overflows int64.
func Add(a, b int64) (int64, bool) {
	// CHANGE CODE BELOW THIS LINE
	if b > 0 && a > math.MaxInt64 {
		// CHANGE CODE ABOVE THIS LINE
		return 0, false
	}
	if b < 0 && a < math.MinInt64-b {
		return 0, false
	}
	return a + b, true
}
