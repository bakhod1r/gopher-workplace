// Package appendcapshared derives a prefix then appends to it. A planted bug
// lets the append clobber the source's tail because they share capacity.
package appendcapshared

// FirstTwoPlus returns the first two elements of xs followed by extra, WITHOUT
// modifying xs. (xs has spare capacity, so a naive sub-slice append overwrites
// xs[2].)
func FirstTwoPlus(xs []int, extra int) []int {
	// CHANGE CODE BELOW THIS LINE
	head := xs[:2]
	// CHANGE CODE ABOVE THIS LINE
	return append(head, extra)
}
