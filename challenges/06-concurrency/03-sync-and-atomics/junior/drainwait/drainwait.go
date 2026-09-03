// Package drainwait — Gopher Workplace challenge.
package drainwait

import "sync"

// Drain tracks in-flight HTTP requests so shutdown can block until the server
// is idle.
type Drain struct {
	mu       sync.Mutex
	cond     *sync.Cond
	inflight int
}

// NewDrain returns a Drain with no requests in flight.
//
// Examples:
//
//	NewDrain().Inflight() => 0
func NewDrain() *Drain {
	// TODO(candidate): build the Drain and attach a Cond to its mutex.
	panic("not implemented")
}

// Start records that a request began.
//
// Examples:
//
//	d := NewDrain(); d.Start(); d.Inflight() => 1
func (d *Drain) Start() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Done records that a request finished, waking any waiter once the server is
// idle.
//
// Examples:
//
//	d.Start(); d.Done(); d.Inflight() => 0
func (d *Drain) Done() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Wait blocks until no requests are in flight.
//
// Examples:
//
//	NewDrain().Wait() => returns immediately
func (d *Drain) Wait() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Inflight returns the number of requests currently running.
func (d *Drain) Inflight() int {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.inflight
}
