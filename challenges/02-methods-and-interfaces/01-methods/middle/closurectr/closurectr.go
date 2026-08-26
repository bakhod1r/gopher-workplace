// Package closurectr — Gopher Workplace challenge.
package closurectr

// Counter tracks a count.
type Counter struct {
	N int
}

// Inc increments by 1 and returns the new count.
func (c *Counter) Inc() int {
	c.N++
	return c.N
}

// NewCounter returns an increment function. Each call to the returned function
// increments and returns the next count, starting from 0.
//
// The key insight: the returned function is a **method value** bound to a
// Counter pointer. Because it's a pointer receiver, all calls share the same
// Counter.
//
// Examples:
//
//	next := NewCounter()
//	next() => 1
//	next() => 2
//	next() => 3
func NewCounter() func() int {
	// TODO(candidate): create a Counter and return its bound Inc method.
	panic("not implemented")
}
