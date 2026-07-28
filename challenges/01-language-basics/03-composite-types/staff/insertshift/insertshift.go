// Package insertshift inserts into a slice by shifting the tail right. A planted
// bug shifts the wrong direction, losing the tail.
package insertshift

// InsertAt inserts v at index i (0..len) in xs, in place, growing by one.
func InsertAt(xs []int, i, v int) []int {
	xs = append(xs, 0) // make room at the end
	// CHANGE CODE BELOW THIS LINE
	copy(xs[i:], xs[i+1:])
	// CHANGE CODE ABOVE THIS LINE
	xs[i] = v
	return xs
}
