// Package rowat — Gopher Workplace challenge.
package rowat

// Row returns the i-th row of g and whether it exists.
//
// An out-of-range index is a missing row, not a panic. The row is returned
// as a view, so writes through it reach g.
//
// Examples:
//
//	Row([][]int{{1}, {2}}, 1) => []int{2}, true
func Row(g [][]int, i int) ([]int, bool) {
	panic("not implemented")
}
