// Package refcount — Gopher Workplace challenge.
package refcount

import "sync"

// Resource is the underlying closeable thing.
type Resource interface {
	Close()
}

// CountingResource records how many times it was closed.
type CountingResource struct {
	mu     sync.Mutex
	Closes int
}

// Close records a close.
func (c *CountingResource) Close() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Closes++
}

// ClosedTimes reports how many times Close was called.
func (c *CountingResource) ClosedTimes() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.Closes
}

// RefCounted owns a resource shared by several holders.
type RefCounted struct {
	mu       sync.Mutex
	res      Resource
	refs     int
	released bool
}

// NewRefCounted returns a handle with one reference held.
func NewRefCounted(r Resource) *RefCounted {
	return &RefCounted{res: r, refs: 1}
}

// Acquire takes another reference. It returns false once the resource has
// been closed.
func (r *RefCounted) Acquire() bool {
	// TODO(candidate): take a reference unless already closed.
	panic("not implemented")
}

// Release drops a reference, closing the resource when the last one goes.
// It returns false when there was nothing to release.
//
// Examples:
//
//	Acquire, Acquire, Release => not closed; count 1
func (r *RefCounted) Release() bool {
	// TODO(candidate): decrement; close exactly once at zero.
	panic("not implemented")
}

// Count returns the current reference count.
func (r *RefCounted) Count() int {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}
