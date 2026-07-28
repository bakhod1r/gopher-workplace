// Package applyinplace transforms the int a pointer references. A planted bug
// calls f but discards its result instead of storing it back through the pointer.
package applyinplace

// Apply replaces *p with f(*p).
func Apply(p *int, f func(int) int) {
	// CHANGE CODE BELOW THIS LINE
	f(*p)
	// CHANGE CODE ABOVE THIS LINE
}
