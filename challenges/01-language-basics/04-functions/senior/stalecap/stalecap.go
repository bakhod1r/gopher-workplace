// Package stalecap builds a running-sum reporter. A planted bug keeps a pointer
// into the slice taken BEFORE an append that reallocates, so later reads see a
// stale backing array.
package stalecap

// FirstAfterGrow appends v to xs (which is full), then returns the value now at
// index 0 of the resulting slice. The append reallocates.
func FirstAfterGrow(v int) int {
	xs := make([]int, 1, 1) // len==cap==1, so append reallocates
	xs[0] = 10
	p := &xs[0]
	xs = append(xs, v)
	// CHANGE CODE BELOW THIS LINE
	*p = 99
	// CHANGE CODE ABOVE THIS LINE
	return xs[0]
}
