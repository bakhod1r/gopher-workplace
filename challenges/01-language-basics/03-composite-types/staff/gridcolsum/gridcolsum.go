// Package gridcolsum sums a column of a grid. A planted bug indexes with the
// row/column roles swapped.
package gridcolsum

// ColSum returns the sum of column c across all rows of a rectangular grid.
func ColSum(grid [][]int, c int) int {
	sum := 0
	for r := range grid {
		// CHANGE CODE BELOW THIS LINE
		sum += grid[c][r]
		// CHANGE CODE ABOVE THIS LINE
	}
	return sum
}
