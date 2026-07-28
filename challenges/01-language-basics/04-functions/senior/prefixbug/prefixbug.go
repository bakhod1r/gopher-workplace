// Package prefixbug answers range-sum queries using a prefix-sum array. A
// planted bug forgets to subtract the prefix at l, over-counting the left part.
package prefixbug

// RangeSum returns the sum of xs[l:r] (l inclusive, r exclusive) using a prefix
// array. Precondition: 0 <= l <= r <= len(xs).
func RangeSum(xs []int, l, r int) int {
	pre := make([]int, len(xs)+1)
	for i, v := range xs {
		pre[i+1] = pre[i] + v
	}
	// CHANGE CODE BELOW THIS LINE
	return pre[r]
	// CHANGE CODE ABOVE THIS LINE
}
