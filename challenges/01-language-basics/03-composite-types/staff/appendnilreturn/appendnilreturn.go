// Package appendnilreturn appends to a slice. A planted bug bails out on a nil
// slice instead of appending (append to nil is valid).
package appendnilreturn

// Add appends x to s and returns the result. A nil s must become a one-element
// slice.
func Add(s []int, x int) []int {
	// CHANGE CODE BELOW THIS LINE
	if s == nil {
		return nil
	}
	// CHANGE CODE ABOVE THIS LINE
	return append(s, x)
}
