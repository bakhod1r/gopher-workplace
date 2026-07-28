// Package subsliceclamp returns a safe sub-slice. A planted bug omits the clamp,
// so an out-of-range end panics.
package subsliceclamp

// Take returns up to n elements from the front of xs. If n exceeds len(xs), it
// returns all of xs (no panic). n <= 0 returns empty.
func Take(xs []int, n int) []int {
	if n <= 0 {
		return xs[:0]
	}
	// CHANGE CODE BELOW THIS LINE
	return xs[:n]
	// CHANGE CODE ABOVE THIS LINE
}
