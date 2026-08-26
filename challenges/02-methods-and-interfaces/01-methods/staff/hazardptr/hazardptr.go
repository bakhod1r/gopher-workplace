// Package hazardptr — Gopher Workplace challenge.
package hazardptr

import "sync/atomic"

// Hazard simulates setting a hazard pointer to protect memory.
type Hazard struct {
	ptr atomic.Pointer[int]
}

// Protect reads the pointer, sets the hazard, and returns the pointer
// if it hasn't changed.
func (h *Hazard) Protect(shared *atomic.Pointer[int]) *int {
	// TODO(candidate):
	// p := shared.Load()
	// h.ptr.Store(p)
	// if shared.Load() == p { return p }
	// return nil
	panic("not implemented")
}
