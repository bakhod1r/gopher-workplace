// Package escapereturn returns a pointer to the max element of a slice. A planted
// bug returns the address of a local COPY (loop value), which aliases nothing in
// the caller's slice, so mutating through it doesn't change the slice.
package escapereturn

// MaxPtr returns a pointer to the largest element of xs (xs non-empty), aliasing
// the slice so writes through it change xs.
func MaxPtr(xs []int) *int {
	best := 0
	for i := range xs {
		if xs[i] > xs[best] {
			best = i
		}
	}
	// CHANGE CODE BELOW THIS LINE
	v := xs[best]
	return &v
	// CHANGE CODE ABOVE THIS LINE
}
