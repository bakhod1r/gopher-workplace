// Package oddbug counts odd numbers. A planted bug tests n%2 == 1, which is
// false for negative odd numbers (e.g. -3 % 2 == -1 in Go).
package oddbug

// CountOdd returns how many elements of xs are odd (including negatives).
func CountOdd(xs []int) int {
	c := 0
	for _, v := range xs {
		// CHANGE CODE BELOW THIS LINE
		if v%2 == 1 {
			// CHANGE CODE ABOVE THIS LINE
			c++
		}
	}
	return c
}
