// Package copyshiftleft shifts a slice left by one (dropping the first element).
// A planted bug uses the wrong copy direction, duplicating data.
package copyshiftleft

// ShiftLeft moves every element one position toward index 0 (xs[0] is lost) and
// sets the last element to 0, in place.
func ShiftLeft(xs []int) {
	if len(xs) == 0 {
		return
	}
	// CHANGE CODE BELOW THIS LINE
	copy(xs[1:], xs)
	// CHANGE CODE ABOVE THIS LINE
	xs[len(xs)-1] = 0
}
