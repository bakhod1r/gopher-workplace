// Package counterregistry — Gopher Workplace challenge.
package counterregistry

import (
	"sync"
	"sync/atomic"
)

// Registry maps metric names to counters. Names are created rarely and
// incremented constantly.
type Registry struct {
	mu       sync.RWMutex
	counters map[string]*atomic.Int64
}

// NewRegistry returns an empty Registry.
//
// Examples:
//
//	NewRegistry().Value("nothing") => 0
func NewRegistry() *Registry {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Add bumps a counter by n, creating it on first use, and returns the new
// value. The common case — an already-registered name — must take only the
// read lock.
//
// Examples:
//
//	r.Add("http_requests", 1) => 1
//	r.Add("http_requests", 2) => 3
func (r *Registry) Add(name string, n int64) int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Value returns a counter's value, or 0 if the name is unknown.
//
// Examples:
//
//	NewRegistry().Value("unknown") => 0
func (r *Registry) Value(name string) int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Snapshot returns a copy of every counter, safe to read after it is returned.
//
// Examples:
//
//	r.Add("a", 2); r.Snapshot() => map[a:2]
func (r *Registry) Snapshot() map[string]int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
