// Package goroutineleak — Gopher Workplace challenge.
package goroutineleak

import "sync"

// Sink receives observed values.
type Sink interface {
	Observe(v int)
}

// CountingSink counts what it observed.
type CountingSink struct {
	mu sync.Mutex
	n  int
}

// Observe records one value.
func (c *CountingSink) Observe(v int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

// Count returns how many values were observed.
func (c *CountingSink) Count() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

// Watcher forwards values from a channel to a sink until told to stop.
type Watcher struct {
	stop     chan struct{}
	stopOnce sync.Once
	done     chan struct{}
}

// NewWatcher returns an unstarted watcher.
func NewWatcher() *Watcher {
	return &Watcher{
		stop: make(chan struct{}),
		done: make(chan struct{}),
	}
}

// Start launches the watching goroutine. The goroutine must exit when in
// closes or Stop is called.
//
// Examples:
//
//	Start then Stop => the goroutine exits
func (w *Watcher) Start(in <-chan int, s Sink) {
	// TODO(candidate): one goroutine with a guaranteed exit path.
	panic("not implemented")
}

// Stop signals the watcher and waits for it to exit. It is idempotent.
func (w *Watcher) Stop() {
	// TODO(candidate): signal once, then wait for the goroutine.
	panic("not implemented")
}

// Wait blocks until the watcher goroutine has exited.
func (w *Watcher) Wait() {
	// TODO(candidate): wait for the done signal.
	panic("not implemented")
}
