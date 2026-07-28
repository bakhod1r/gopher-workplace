// Package factbug computes factorial recursively. A planted bug uses the wrong
// base case (returns 0 at n==0), zeroing every result.
package factbug

// Fact returns n! for n >= 0.
func Fact(n int) int {
	if n == 0 {
		// CHANGE CODE BELOW THIS LINE
		return 0
		// CHANGE CODE ABOVE THIS LINE
	}
	return n * Fact(n-1)
}
