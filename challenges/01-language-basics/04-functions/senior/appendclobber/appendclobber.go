// Package appendclobber derives two slices from a shared base with spare
// capacity. A planted bug appends to the base directly, so the second append
// overwrites the first result's tail.
package appendclobber

// TwoTails returns base+ x and base+ y as two independent slices. base may have
// spare capacity, so appending to it directly would make the two results share
// (and clobber) memory.
func TwoTails(base []int, x, y int) (a, b []int) {
	// CHANGE CODE BELOW THIS LINE
	a = append(base, x)
	b = append(base, y)
	// CHANGE CODE ABOVE THIS LINE
	return a, b
}
