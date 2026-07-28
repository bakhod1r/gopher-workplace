// Package sortmutates returns a sorted copy. A planted bug sorts the input in
// place, corrupting the caller's slice.
package sortmutates

import "sort"

// SortedCopy returns a new sorted slice; the input xs must not be modified.
func SortedCopy(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	out := xs
	// CHANGE CODE ABOVE THIS LINE
	sort.Ints(out)
	return out
}
