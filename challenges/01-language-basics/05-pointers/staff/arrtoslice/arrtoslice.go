// Package arrtoslice converts a *[4]int into a full []int view. A planted bug
// slices with a fixed wrong length, dropping elements.
package arrtoslice

// AsSlice returns a slice viewing all elements of the array p points to.
func AsSlice(p *[4]int) []int {
	// CHANGE CODE BELOW THIS LINE
	return p[:2]
	// CHANGE CODE ABOVE THIS LINE
}
