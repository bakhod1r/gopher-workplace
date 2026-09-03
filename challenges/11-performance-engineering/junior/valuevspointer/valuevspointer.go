// Package valuevspointer — Gopher Workplace challenge.
package valuevspointer

// Counter tracks a running total plus a fixed-size history buffer, so a copy
// of it is expensive rather than free.
type Counter struct {
	n       int
	history [64]int
}

// Inc adds one to the counter. It must mutate the receiver the caller holds.
//
// Examples:
//
//	c.Inc(); c.Value() => 1
func (c *Counter) Inc() {
	panic("not implemented")
}

// Value returns the current count without copying the whole struct.
//
// Examples:
//
//	(&Counter{}).Value() => 0
func (c *Counter) Value() int {
	panic("not implemented")
}

// IncCopy takes the receiver by value, so the increment is lost. It returns
// the incremented copy, which is the only way to observe the change.
//
// Examples:
//
//	c.IncCopy().Value() => 1, while c.Value() is still 0
func (c Counter) IncCopy() Counter {
	panic("not implemented")
}
