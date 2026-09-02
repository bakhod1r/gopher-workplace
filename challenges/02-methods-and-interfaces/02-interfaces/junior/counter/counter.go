// Package counter — Gopher Workplace challenge.
package counter

// Counter reports a running total.
type Counter interface {
	Count() int
}

// Clicks accumulates click events.
type Clicks struct {
	N int
}

// Count returns the number of clicks.
func (c *Clicks) Count() int {
	// TODO(candidate): return the stored count.
	panic("not implemented")
}

// Fixed is a constant counter.
type Fixed int

// Count returns the constant value.
func (f Fixed) Count() int {
	// TODO(candidate): return the value itself.
	panic("not implemented")
}

// Total sums every counter.
//
// Examples:
//
//	Total([]Counter{&Clicks{N: 3}, Fixed(2)}) => 5
func Total(cs []Counter) int {
	// TODO(candidate): add up Count() over the slice.
	panic("not implemented")
}
