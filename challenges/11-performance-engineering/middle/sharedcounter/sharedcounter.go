// Package sharedcounter — Gopher Workplace challenge.
package sharedcounter

import "sync"

// Counter is a set of named counters safe for concurrent use. The zero value
// is ready to use.
type Counter struct {
	mu sync.Mutex
	n  map[string]int64
}

// Add increases a counter by delta. Concurrent Adds must not lose updates:
// read-modify-write on shared state is three operations, and another
// goroutine can land between any two of them.
//
// Examples:
//
//	c.Add("hits", 1)
func (c *Counter) Add(key string, delta int64) {
	panic("not implemented")
}

// Get returns one counter's value.
//
// Examples:
//
//	c.Get("hits") => 1
func (c *Counter) Get(key string) int64 {
	panic("not implemented")
}

// Snapshot returns a copy of every counter, so the caller can range it
// without holding the lock — and without racing the next Add.
//
// Examples:
//
//	c.Snapshot() => map[string]int64{"hits": 1}
func (c *Counter) Snapshot() map[string]int64 {
	panic("not implemented")
}
