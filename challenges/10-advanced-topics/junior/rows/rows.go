// Package rows — Gopher Workplace challenge.
package rows

// Rows returns an r-by-c grid of zeros whose rows are consecutive
// windows into a single backing array.
//
// Allocating each row separately costs r allocations and scatters the grid
// across the heap; this must cost two.
//
// Examples:
//
//	Rows(2, 3) => a 2x3 grid, rows 0 and 1 adjacent in memory
func Rows(r, c int) [][]int {
	panic("not implemented")
}
