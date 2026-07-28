// Package rangemutate doubles every element in place. A planted bug writes to
// the range VALUE variable, which is a copy, so the slice is unchanged.
package rangemutate

// DoubleAll doubles each element of xs in place.
func DoubleAll(xs []int) {
	for i, v := range xs {
		_ = i
		// CHANGE CODE BELOW THIS LINE
		v = v * 2
		_ = v
		// CHANGE CODE ABOVE THIS LINE
	}
}
