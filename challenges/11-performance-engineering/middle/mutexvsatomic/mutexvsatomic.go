// Package mutexvsatomic — Gopher Workplace challenge.
package mutexvsatomic

import (
	"sync"
	"sync/atomic"
)

// MutexCounter guards a plain int with a mutex.
type MutexCounter struct {
	mu sync.Mutex
	n  int64
}

// Inc adds one.
//
// Examples:
//
//	c.Inc()
func (c *MutexCounter) Inc() {
	panic("not implemented")
}

// Value returns the count.
//
// Examples:
//
//	c.Value() => 1
func (c *MutexCounter) Value() int64 {
	panic("not implemented")
}

// AtomicCounter does the same job with one atomic word.
type AtomicCounter struct {
	n atomic.Int64
}

// Inc adds one.
//
// Examples:
//
//	c.Inc()
func (c *AtomicCounter) Inc() {
	panic("not implemented")
}

// Value returns the count.
//
// Examples:
//
//	c.Value() => 1
func (c *AtomicCounter) Value() int64 {
	panic("not implemented")
}

// ShardedCounter avoids the shared word entirely: each shard is its own
// counter, and the total is their sum. This is the version that actually
// scales, because two cores incrementing different shards never touch the
// same cache line's worth of contention.
type ShardedCounter struct {
	shards []atomic.Int64
}

// NewSharded returns a counter with n shards; a non-positive n gives 1.
//
// Examples:
//
//	NewSharded(8)
func NewSharded(n int) *ShardedCounter {
	panic("not implemented")
}

// Inc adds one to the shard chosen by id, so callers with different ids
// mostly touch different shards.
//
// Examples:
//
//	c.Inc(7)
func (c *ShardedCounter) Inc(id int) {
	panic("not implemented")
}

// Value returns the sum of every shard. It is only approximate under
// concurrent writers — shards read at slightly different instants — which is
// the trade the design makes.
//
// Examples:
//
//	c.Value() => 1
func (c *ShardedCounter) Value() int64 {
	panic("not implemented")
}
