// Package nilmap tallies counts. A planted bug writes to a nil map.
package nilmap

// Count returns a map from each value to its occurrence count.
func Count(xs []int) map[int]int {
	// CHANGE CODE BELOW THIS LINE
	var m map[int]int
	// CHANGE CODE ABOVE THIS LINE
	for _, x := range xs {
		m[x]++
	}
	return m
}
