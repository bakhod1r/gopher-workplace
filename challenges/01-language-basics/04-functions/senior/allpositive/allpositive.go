// Package allpositive reports whether every element is positive. A planted bug
// returns true from inside the loop on the first positive element, so it never
// checks the rest.
package allpositive

// AllPositive reports whether all elements of xs are > 0. Empty is true.
func AllPositive(xs []int) bool {
	for _, v := range xs {
		// CHANGE CODE BELOW THIS LINE
		if v > 0 {
			return true
		}
		// CHANGE CODE ABOVE THIS LINE
	}
	return true
}
