// Package labeledcont counts rows of a grid that contain NO negative number. A
// planted bug uses a bare continue, which only skips to the next column instead
// of abandoning the whole row, so rows with negatives are still counted.
package labeledcont

// CleanRows returns how many rows contain no negative value.
func CleanRows(grid [][]int) int {
	count := 0
Rows:
	for _, row := range grid {
		for _, v := range row {
			if v < 0 {
				// CHANGE CODE BELOW THIS LINE
				break Rows
				// CHANGE CODE ABOVE THIS LINE
			}
		}
		count++
	}
	return count
}
