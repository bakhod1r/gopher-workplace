// Package grid2dalias builds a 2-D grid. A planted bug makes all rows share one
// backing slice.
package grid2dalias

// New builds a rows x cols grid of zeros. Writing one cell must not affect
// others in different rows.
func New(rows, cols int) [][]int {
	grid := make([][]int, rows)
	// CHANGE CODE BELOW THIS LINE
	row := make([]int, cols)
	for i := range grid {
		grid[i] = row
	}
	// CHANGE CODE ABOVE THIS LINE
	return grid
}
