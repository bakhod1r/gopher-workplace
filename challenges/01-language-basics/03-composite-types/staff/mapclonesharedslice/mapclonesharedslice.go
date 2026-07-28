// Package mapclonesharedslice deep-copies a map of slices. A planted bug shares
// the slice values, so mutating the clone corrupts the original.
package mapclonesharedslice

// Clone returns an independent copy: both the map and each slice value must be
// separate from the input.
func Clone(m map[string][]int) map[string][]int {
	out := make(map[string][]int, len(m))
	for k, v := range m {
		// CHANGE CODE BELOW THIS LINE
		out[k] = v
		// CHANGE CODE ABOVE THIS LINE
	}
	return out
}
