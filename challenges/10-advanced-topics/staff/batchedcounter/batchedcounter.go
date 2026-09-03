// Package batchedcounter — Gopher Workplace challenge.
package batchedcounter

import "sync/atomic"

// batchSize is how much a Local may accumulate before publishing.
const batchSize = 64

// Local is one goroutine's private accumulator. It must not be shared.
type Local struct {
	n int64
}

// Counter is a shared total fed by batched local accumulators.
type Counter struct {
	total atomic.Int64
}

// Flush publishes whatever the local still holds.
func (c *Counter) Flush(local *Local) {
	if local.n != 0 {
		c.total.Add(local.n)
		local.n = 0
	}
}

// Total returns the published total.
func (c *Counter) Total() int64 { return c.total.Load() }

// Add adds n to the caller's local accumulator, flushing it into the
// shared total when it reaches the batch threshold.
//
// The shared atomic is the contended resource; touching it once per batch
// instead of once per event is the whole point.
//
// Examples:
//
//	c.Add(local, 1) a thousand times => the total is 1000 after Flush
func (c *Counter) Add(local *Local, n int64) {
	panic("not implemented")
}
