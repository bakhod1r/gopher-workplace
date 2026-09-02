// Package ifaceembed — Gopher Workplace challenge.
package ifaceembed

// Source reports a number.
type Source interface {
	Value() int
}

// Gauge is a fixed reading.
type Gauge struct {
	N int
}

// Value returns the reading.
func (g Gauge) Value() int {
	// TODO(candidate): return N.
	panic("not implemented")
}

// CountingSource wraps any Source and counts reads.
type CountingSource struct {
	Source
	Calls int
}

// Value counts the read and delegates to the wrapped source.
//
// Examples:
//
//	c := &CountingSource{Source: Gauge{N: 5}}
//	c.Value() => 5, and c.Calls == 1
func (c *CountingSource) Value() int {
	// TODO(candidate): increment Calls, delegate to the embedded Source.
	panic("not implemented")
}

// ReadTwice reads s twice and returns both readings.
func ReadTwice(s Source) (int, int) {
	// TODO(candidate): two reads, in order.
	panic("not implemented")
}
