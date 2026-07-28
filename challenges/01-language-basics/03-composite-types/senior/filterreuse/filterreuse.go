// Package filterreuse returns even numbers without modifying the input. A
// planted bug filters in place over the same backing array, corrupting xs.
package filterreuse

// Evens returns a new slice of the even elements of xs. xs must not be modified.
func Evens(xs []int) []int {
	// CHANGE CODE BELOW THIS LINE
	out := xs[:0]
	// CHANGE CODE ABOVE THIS LINE
	for _, x := range xs {
		if x%2 == 0 {
			out = append(out, x)
		}
	}
	return out
}
