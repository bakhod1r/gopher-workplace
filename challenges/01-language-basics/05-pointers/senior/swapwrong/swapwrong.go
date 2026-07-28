// Package swapwrong swaps two ints through pointers. A planted bug swaps the
// local pointer variables instead of the pointed-to values, so the caller sees
// no change.
package swapwrong

// Swap should exchange the values that a and b point to.
func Swap(a, b *int) {
	// CHANGE CODE BELOW THIS LINE
	a, b = b, a
	// CHANGE CODE ABOVE THIS LINE
}
