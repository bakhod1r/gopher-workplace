// Package removeorder deletes a slice element by index. A planted off-by-one
// keeps the removed element.
package removeorder

// RemoveAt returns xs with the element at index i removed (0 <= i < len).
func RemoveAt(xs []int, i int) []int {
	// CHANGE CODE BELOW THIS LINE
	return append(xs[:i], xs[i:]...)
	// CHANGE CODE ABOVE THIS LINE
}
