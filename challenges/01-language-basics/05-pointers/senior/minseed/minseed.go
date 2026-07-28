// Package minseed returns a pointer to the smallest element. A planted bug seeds
// the min with the last element's address but compares from index 0, or rather
// seeds with a nil/zero assumption; here it starts min at &xs[0] but the loop
// skips index 0 incorrectly by starting comparisons that overwrite wrongly. The
// concrete bug: it seeds min = &xs[len-1] instead of &xs[0].
package minseed

// MinPtr returns a pointer to the smallest element of xs (xs non-empty). On a
// tie it returns the earliest.
func MinPtr(xs []int) *int {
	// CHANGE CODE BELOW THIS LINE
	min := &xs[len(xs)-1]
	// CHANGE CODE ABOVE THIS LINE
	for i := range xs {
		if xs[i] < *min {
			min = &xs[i]
		}
	}
	return min
}
