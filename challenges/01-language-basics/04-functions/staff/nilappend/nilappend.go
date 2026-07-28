// Package nilappend collects results into a slice that starts nil. A planted bug
// bails out early when the slice is nil, thinking it must be initialised first,
// so it returns nothing.
package nilappend

// Collect returns xs with each element of extra appended. xs may be nil, which
// is a valid empty slice to append to.
func Collect(xs []int, extra []int) []int {
	// CHANGE CODE BELOW THIS LINE
	if xs == nil {
		return nil
	}
	// CHANGE CODE ABOVE THIS LINE
	for _, v := range extra {
		xs = append(xs, v)
	}
	return xs
}
