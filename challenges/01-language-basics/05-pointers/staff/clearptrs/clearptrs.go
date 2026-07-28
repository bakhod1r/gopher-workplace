// Package clearptrs clears a slice of pointers so referenced objects can be
// freed. A planted bug re-slices to length 0 but leaves the pointers in the
// backing array, retaining the objects.
package clearptrs

// Clear empties the slice AND drops its references so the pointees can be GC'd.
// It returns the emptied slice (len 0).
func Clear(s []*int) []*int {
	// CHANGE CODE BELOW THIS LINE
	return s[:0]
	// CHANGE CODE ABOVE THIS LINE
}
