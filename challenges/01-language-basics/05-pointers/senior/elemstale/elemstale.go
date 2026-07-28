// Package elemstale returns a pointer into a slice, then appends. A planted bug
// keeps the pointer taken before a reallocating append, so writes through it
// don't affect the returned slice.
package elemstale

// FirstOf returns a slice built from xs plus a trailing 0, with its first
// element set to 42 through a pointer.
func FirstOf(xs []int) []int {
	s := make([]int, len(xs), len(xs)) // len == cap, so append reallocates
	copy(s, xs)
	p := &s[0]
	s = append(s, 0)
	// CHANGE CODE BELOW THIS LINE
	*p = 42
	// CHANGE CODE ABOVE THIS LINE
	return s
}
