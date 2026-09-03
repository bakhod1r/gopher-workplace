// Package poolwrapper — Gopher Workplace challenge.
package poolwrapper

import "sync"

// Pool hands out reusable byte buffers. The zero value is ready to use, and
// it is safe for concurrent use because sync.Pool is.
type Pool struct {
	pool sync.Pool
	Size int // capacity of a freshly created buffer; 0 means 1024
}

// Get returns a buffer with length 0 and capacity at least the pool's size.
// A buffer that has been through Put comes back empty, never carrying the
// previous caller's bytes.
//
// Examples:
//
//	b := p.Get(); len(b) == 0
func (p *Pool) Get() []byte {
	panic("not implemented")
}

// Put returns a buffer to the pool for reuse. A nil buffer is ignored.
//
// Examples:
//
//	p.Put(b)
func (p *Pool) Put(b []byte) {
	panic("not implemented")
}
