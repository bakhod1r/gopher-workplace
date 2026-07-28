// Package elemptr returns a pointer into a slice and the grown slice. A planted
// bug takes the element pointer, then grows the slice (reallocating), so the
// pointer is stale and writing through it doesn't affect the returned slice.
package elemptr

// BumpHead appends v (reallocating a full slice), then must make the first
// element of the RESULT equal to 42 by writing through a pointer.
func BumpHead(v int) []int {
	xs := make([]int, 1, 1)
	xs[0] = 1
	p := &xs[0]
	xs = append(xs, v)
	// CHANGE CODE BELOW THIS LINE
	*p = 42
	// CHANGE CODE ABOVE THIS LINE
	return xs
}
