// Package prealloc builds a slice of doubles. A planted bug preallocates with a
// LENGTH of n (not capacity), then appends, producing n leading zeros followed
// by the real data.
package prealloc

// Doubles returns [2,4,6,...,2n].
func Doubles(n int) []int {
	// CHANGE CODE BELOW THIS LINE
	out := make([]int, n)
	// CHANGE CODE ABOVE THIS LINE
	for i := 1; i <= n; i++ {
		out = append(out, i*2)
	}
	return out
}
