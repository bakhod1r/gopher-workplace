// Package accreset sums a slice with a tail-recursive helper threading an
// accumulator. A planted bug ignores the passed accumulator and restarts from
// zero at each call, so only the last element survives.
package accreset

func sumAcc(xs []int, acc int) int {
	if len(xs) == 0 {
		return acc
	}
	// CHANGE CODE BELOW THIS LINE
	return sumAcc(xs[1:], xs[0])
	// CHANGE CODE ABOVE THIS LINE
}

// Sum returns the total of xs.
func Sum(xs []int) int {
	return sumAcc(xs, 0)
}
