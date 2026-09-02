// Package atomicvsmutex — Gopher Workplace challenge.
package atomicvsmutex

import (
	"sync"
	"sync/atomic"
)

// Counter is a monotonically increasing counter.
type Counter interface {
	Inc()
	Add(delta int64)
	Value() int64
}

// MutexCounter guards its value with a mutex.
type MutexCounter struct {
	mu sync.Mutex
	n  int64
}

// Inc adds one.
func (c *MutexCounter) Inc() {
	// TODO(candidate): increment under the lock.
	panic("not implemented")
}

// Add adds delta.
func (c *MutexCounter) Add(delta int64) {
	// TODO(candidate): add under the lock.
	panic("not implemented")
}

// Value returns the current count.
func (c *MutexCounter) Value() int64 {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}

// AtomicCounter uses an atomic integer.
type AtomicCounter struct {
	n atomic.Int64
}

// Inc adds one.
func (c *AtomicCounter) Inc() {
	// TODO(candidate): atomic add.
	panic("not implemented")
}

// Add adds delta.
func (c *AtomicCounter) Add(delta int64) {
	// TODO(candidate): atomic add.
	panic("not implemented")
}

// Value returns the current count.
func (c *AtomicCounter) Value() int64 {
	// TODO(candidate): atomic load.
	panic("not implemented")
}

// IncAll increments every counter once.
//
// Examples:
//
//	IncAll(a, b) => both are incremented
func IncAll(cs ...Counter) {
	// TODO(candidate): increment each counter.
	panic("not implemented")
}
