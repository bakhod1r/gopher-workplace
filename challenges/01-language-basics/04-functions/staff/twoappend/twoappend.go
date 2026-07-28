// Package twoappend forks a buffer into two variants. A planted bug appends to
// the same base with spare capacity twice, so the second append overwrites the
// first variant's last element.
package twoappend

// Fork returns base+a and base+b as two independent slices.
func Fork(base []int, a, b int) ([]int, int) {
	// CHANGE CODE BELOW THIS LINE
	x := append(base, a)
	y := append(base, b)
	// CHANGE CODE ABOVE THIS LINE
	_ = y
	return x, x[len(x)-1]
}
