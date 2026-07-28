// Package aliasbug returns a modified copy of a slice. A planted bug assigns the
// slice (copying only the header), so mutating the "copy" corrupts the input.
package aliasbug

// WithFirst returns a copy of xs whose first element is v, leaving xs unchanged.
// Precondition: len(xs) > 0.
func WithFirst(xs []int, v int) []int {
	// CHANGE CODE BELOW THIS LINE
	cp := xs
	// CHANGE CODE ABOVE THIS LINE
	cp[0] = v
	return cp
}
