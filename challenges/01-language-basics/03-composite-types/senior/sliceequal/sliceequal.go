// Package sliceequal compares two int slices. A planted bug compares only
// lengths.
package sliceequal

// Equal reports whether a and b have the same length and equal elements in
// order.
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	// CHANGE CODE BELOW THIS LINE
	return true
	// CHANGE CODE ABOVE THIS LINE
}
