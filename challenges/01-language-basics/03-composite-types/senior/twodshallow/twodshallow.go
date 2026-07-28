// Package twodshallow clones a 2-D grid. A planted bug copies only the outer
// slice, so rows stay shared.
package twodshallow

// Clone returns an independent copy of a grid (rows must not be shared).
func Clone(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	// CHANGE CODE BELOW THIS LINE
	copy(out, grid)
	// CHANGE CODE ABOVE THIS LINE
	return out
}
