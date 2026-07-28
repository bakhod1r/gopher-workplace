// Package appendwhilerange duplicates each element into the SAME slice. A
// planted bug ranges over the live slice while appending to it; range snapshots
// the ORIGINAL length, but the bug instead loops on a manual index bound that
// grows, never terminating conceptually — here it reads grown data and doubles
// wrongly. Fix: range a snapshot / fixed length.
package appendwhilerange

// DupAll returns a new slice containing each element of xs followed by its
// double: [a, 2a, b, 2b, ...].
func DupAll(xs []int) []int {
	out := make([]int, 0, len(xs)*2)
	// CHANGE CODE BELOW THIS LINE
	for i := 0; i < len(out); i++ {
		out = append(out, xs[i], xs[i]*2)
	}
	// CHANGE CODE ABOVE THIS LINE
	return out
}
