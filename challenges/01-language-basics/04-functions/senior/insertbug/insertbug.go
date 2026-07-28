// Package insertbug inserts a value at an index, shifting the tail right. A
// planted bug copies in the wrong direction, duplicating an element and losing
// the tail.
package insertbug

// InsertAt returns xs with v inserted at index i (0 <= i <= len(xs)).
func InsertAt(xs []int, i, v int) []int {
	xs = append(xs, 0)
	// CHANGE CODE BELOW THIS LINE
	copy(xs[i:], xs[i+1:])
	// CHANGE CODE ABOVE THIS LINE
	xs[i] = v
	return xs
}
