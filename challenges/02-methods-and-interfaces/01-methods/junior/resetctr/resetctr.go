// Package resetctr — Gopher Workplace challenge.
package resetctr

// Counter holds a count.
type Counter struct {
	N int
}

// Reset sets the counter back to zero.
//
// Examples:
//
//	c := Counter{42}; c.Reset() // c.N == 0
func (c *Counter) Reset() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.N
}
