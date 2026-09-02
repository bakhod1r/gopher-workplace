// Package counterstruct — Gopher Workplace challenge.
package counterstruct

// Counter tallies how often each value of T is seen.
// Use NewCounter to create one.
type Counter[T comparable] struct {
	counts map[T]int
}

// NewCounter returns a ready-to-use counter.
func NewCounter[T comparable]() *Counter[T] {
	// TODO(candidate): return a counter with its map allocated.
	panic("not implemented")
}

// Inc adds one to the tally for v.
func (c *Counter[T]) Inc(v T) {
	// TODO(candidate): increment the tally for v.
	panic("not implemented")
}

// Count returns the tally for v, or 0 when v was never seen.
func (c *Counter[T]) Count(v T) int {
	// TODO(candidate): report the tally for v.
	panic("not implemented")
}

// Total returns the sum of all tallies.
func (c *Counter[T]) Total() int {
	// TODO(candidate): sum every tally.
	panic("not implemented")
}
