// Package compactcopy removes the element at index 0 by shifting left with copy.
// A planted bug copies in the wrong direction, so the shift duplicates rather
// than compacts.
package compactcopy

// DropFirst removes the first element by shifting the tail left, then trims the
// length. Returns the compacted slice.
func DropFirst(xs []int) []int {
	if len(xs) == 0 {
		return xs
	}
	// CHANGE CODE BELOW THIS LINE
	copy(xs[1:], xs[:len(xs)-1])
	// CHANGE CODE ABOVE THIS LINE
	return xs[:len(xs)-1]
}
