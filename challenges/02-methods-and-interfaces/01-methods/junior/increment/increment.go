// Package increment — Gopher Workplace challenge.
package increment

// Counter holds a count.
type Counter struct {
	N int
}

// Inc adds 1 to the counter.
//
// Examples:
//
//	c := Counter{0}; c.Inc() // c.N == 1
//	c := Counter{5}; c.Inc() // c.N == 6
func (c *Counter) Inc() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Value returns the current count.
func (c Counter) Value() int {
	return c.N
}
