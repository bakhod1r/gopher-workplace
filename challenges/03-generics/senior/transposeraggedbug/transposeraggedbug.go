// Package transposeraggedbug — Gopher Workplace challenge.
package transposeraggedbug

// Transpose swaps rows and columns of a possibly ragged matrix.
// Column i of the result holds every row that is long enough.
//
// Examples:
//
//	Transpose([][]int{{1, 2}, {3}}) => [][]int{{1, 3}, {2}}
func Transpose[T any](m [][]T) [][]T {
	// CHANGE CODE BELOW THIS LINE
	width := 0
	if len(m) > 0 {
		width = len(m[0])
	}
	out := make([][]T, width)
	for c := 0; c < width; c++ {
		col := make([]T, 0, len(m))
		for _, row := range m {
			if c < len(row) {
				col = append(col, row[c])
			}
		}
		out[c] = col
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
