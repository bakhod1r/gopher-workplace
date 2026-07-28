// Package maparraykey counts grid cells using an array key. A planted bug swaps
// the coordinate order in the key.
package maparraykey

// CountCells tallies how many times each (row, col) coordinate appears, keyed by
// a [2]int array.
func CountCells(cells [][2]int) map[[2]int]int {
	m := make(map[[2]int]int)
	for _, c := range cells {
		// CHANGE CODE BELOW THIS LINE
		m[[2]int{c[1], c[0]}]++
		// CHANGE CODE ABOVE THIS LINE
	}
	return m
}
