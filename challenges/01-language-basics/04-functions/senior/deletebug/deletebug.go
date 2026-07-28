// Package deletebug removes the element at index i, preserving order. A planted
// bug shifts the wrong direction, dropping the wrong element.
package deletebug

// RemoveAt returns xs with the element at index i removed, order preserved.
// Precondition: 0 <= i < len(xs).
func RemoveAt(xs []int, i int) []int {
	// CHANGE CODE BELOW THIS LINE
	return append(xs[:i], xs[i:]...)
	// CHANGE CODE ABOVE THIS LINE
}
