// Package uploadwindow — Gopher Workplace challenge.
package uploadwindow

import "sync"

// Window limits how many uploads may run at the same time.
type Window struct {
	mu    sync.Mutex
	cond  *sync.Cond
	limit int
	inUse int
}

// NewWindow returns a Window allowing limit concurrent uploads. A limit of
// zero or less collapses to one.
//
// Examples:
//
//	NewWindow(3).InUse() => 0
func NewWindow(limit int) *Window {
	// TODO(candidate): build the Window and attach a Cond to its mutex.
	panic("not implemented")
}

// Acquire takes a permit, blocking while the window is full.
//
// Examples:
//
//	w := NewWindow(1); w.Acquire(); w.InUse() => 1
func (w *Window) Acquire() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Release returns a permit and wakes any goroutine waiting for one.
//
// Examples:
//
//	w.Acquire(); w.Release(); w.InUse() => 0
func (w *Window) Release() {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// InUse reports how many permits are held right now.
func (w *Window) InUse() int {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.inUse
}
