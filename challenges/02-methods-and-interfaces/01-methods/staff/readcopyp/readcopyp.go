// Package readcopyp — Gopher Workplace challenge.
package readcopyp

import (
	"sync"
	"sync/atomic"
)

// Config holds read-mostly state.
type Config struct {
	Data string
}

// RCU implements Read-Copy-Update semantics.
type RCU struct {
	ptr atomic.Pointer[Config]
	mu  sync.Mutex // only for writers
}

func New() *RCU {
	r := &RCU{}
	r.ptr.Store(&Config{Data: "v1"})
	return r
}

// Update creates a copy, modifies it, and swaps the pointer.
func (r *RCU) Update(newData string) {
	// TODO(candidate): Lock mu. Load ptr. Create new copy with newData. Store pointer. Unlock.
	panic("not implemented")
}
