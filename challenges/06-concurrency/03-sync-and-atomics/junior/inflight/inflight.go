// Package inflight - Gopher Workplace challenge.
package inflight

import "sync/atomic"

// Gauge tracks how many requests are currently in flight.
type Gauge struct {
	n atomic.Int64
}

// Enter records that a request started.
//
// Examples:
//
//	var g Gauge; g.Enter(); g.Current()            => 1
//	var g Gauge; g.Enter(); g.Enter(); g.Current() => 2
func (g *Gauge) Enter() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Exit records that a request finished.
//
// Examples:
//
//	var g Gauge; g.Enter(); g.Exit(); g.Current() => 0
func (g *Gauge) Exit() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Current returns the number of requests in flight.
func (g *Gauge) Current() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
