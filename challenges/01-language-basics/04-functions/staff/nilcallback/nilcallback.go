// Package nilcallback invokes an optional hook. A planted bug calls the hook
// without checking for nil, so a nil hook panics instead of being skipped.
package nilcallback

// Process applies hook to each element if hook is non-nil, summing the results;
// if hook is nil it sums the elements unchanged.
func Process(xs []int, hook func(int) int) int {
	total := 0
	for _, v := range xs {
		// CHANGE CODE BELOW THIS LINE
		total += hook(v)
		// CHANGE CODE ABOVE THIS LINE
	}
	return total
}
