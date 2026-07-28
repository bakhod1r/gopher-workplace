// Package arrayvalue doubles every element of a fixed array through a pointer.
// A planted bug works on a copy of the array, so the caller's array is untouched.
package arrayvalue

// Double doubles each element of the pointed-to 3-element array in place.
func Double(a *[3]int) {
	// CHANGE CODE BELOW THIS LINE
	b := *a
	for i := range b {
		b[i] *= 2
	}
	// CHANGE CODE ABOVE THIS LINE
}
