// Package retainbacking extracts a small prefix to keep long-term. A planted bug
// returns a plain re-slice, which keeps the whole backing array alive AND lets
// an append spill into the parent. The test detects the spill.
package retainbacking

// Prefix returns the first k elements as a slice whose appends cannot reach the
// parent's remaining elements.
func Prefix(xs []int, k int) []int {
	// CHANGE CODE BELOW THIS LINE
	return xs[:k]
	// CHANGE CODE ABOVE THIS LINE
}
