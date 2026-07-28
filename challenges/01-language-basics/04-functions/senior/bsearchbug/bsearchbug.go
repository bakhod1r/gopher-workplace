// Package bsearchbug finds a target in a sorted slice. A planted bug moves the
// low pointer to mid (not mid+1), so on a miss the search can loop or overshoot;
// the tests catch a wrong/absent index.
package bsearchbug

// IndexOf returns the index of target in the sorted slice xs, or -1 if absent.
func IndexOf(xs []int, target int) int {
	lo, hi := 0, len(xs)-1
	for lo <= hi {
		mid := (lo + hi) / 2
		switch {
		case xs[mid] == target:
			return mid
		case xs[mid] < target:
			// CHANGE CODE BELOW THIS LINE
			lo = mid
			// CHANGE CODE ABOVE THIS LINE
		default:
			hi = mid - 1
		}
	}
	return -1
}
