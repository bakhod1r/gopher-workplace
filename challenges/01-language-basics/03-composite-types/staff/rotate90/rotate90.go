// Package rotate90 rotates a square matrix 90 degrees clockwise. A planted bug
// forgets to reverse rows after transposing.
package rotate90

// Rotate returns an n×n matrix rotated 90° clockwise. Rotation is transpose,
// then reverse each row.
func Rotate(m [][]int) [][]int {
	n := len(m)
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			out[i][j] = m[j][i] // transpose
		}
	}
	// CHANGE CODE BELOW THIS LINE
	// (reverse each row of out here)
	// CHANGE CODE ABOVE THIS LINE
	return out
}
