// Package growthstale demonstrates that append can reallocate, invalidating a
// pointer into the old backing array. A planted bug writes through the stale
// pointer.
package growthstale

// BumpFirst appends val to a full-capacity slice (forcing reallocation) and then
// sets the first element to 99. Returns the first element and the grown slice.
func BumpFirst(val int) (int, []int) {
	s := make([]int, 1, 1) // len 1, cap 1 (full)
	s[0] = 10
	p := &s[0]
	s = append(s, val) // reallocates; p now points at the OLD array
	// CHANGE CODE BELOW THIS LINE
	*p = 99
	// CHANGE CODE ABOVE THIS LINE
	return s[0], s
}
