// Package reversebug reverses a slice in place with two pointers. A planted bug
// initialises the right pointer to len(xs) instead of len(xs)-1, so the first
// access xs[j] indexes one past the end and panics.
package reversebug

// Reverse reverses xs in place.
func Reverse(xs []int) {
	// CHANGE CODE BELOW THIS LINE
	i, j := 0, len(xs)
	// CHANGE CODE ABOVE THIS LINE
	for i < j {
		xs[i], xs[j] = xs[j], xs[i]
		i++
		j--
	}
}
