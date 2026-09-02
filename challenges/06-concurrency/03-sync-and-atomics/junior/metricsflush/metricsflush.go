// Package metricsflush - Gopher Workplace challenge.
package metricsflush

import "sync/atomic"

// Meter accumulates a byte counter that is flushed periodically.
type Meter struct {
	n atomic.Int64
}

// Record adds v to the pending total.
//
// Examples:
//
//	var m Meter; m.Record(3); m.Record(4); m.Pending() => 7
func (m *Meter) Record(v int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Drain returns the pending total and resets it to zero.
//
// Examples:
//
//	var m Meter; m.Record(3); m.Drain() => 3
//	var m Meter; m.Drain()              => 0
func (m *Meter) Drain() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Pending returns the current total without clearing it.
func (m *Meter) Pending() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
