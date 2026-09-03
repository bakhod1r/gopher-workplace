// Package ringpool — Gopher Workplace challenge.
package ringpool

import "sync/atomic"

// BufPool hands out fixed-size buffers from a bounded ring.
type BufPool struct {
	size  int
	next  atomic.Int64
	slots []atomic.Pointer[[]byte]
}

// NewBufPool returns a pool of n slots holding size-byte buffers.
func NewBufPool(n, size int) *BufPool {
	if n < 1 {
		n = 1
	}
	if size < 1 {
		size = 1
	}
	return &BufPool{size: size, slots: make([]atomic.Pointer[[]byte], n)}
}

// Put returns a buffer to the ring, dropping it if the ring is full or the
// buffer is the wrong size.
func (p *BufPool) Put(b []byte) {
	if cap(b) != p.size {
		return
	}
	b = b[:0]
	i := int(p.next.Add(1)-1) % len(p.slots)
	if i < 0 {
		i = -i
	}
	p.slots[i].CompareAndSwap(nil, &b)
}

// Get returns a buffer from the ring, or a fresh one when the ring is
// empty.
//
// The ring is a fixed array of slots claimed with an atomic index, so
// concurrent callers never block each other and never receive the same
// buffer twice.
//
// Examples:
//
//	p := NewBufPool(4, 64); p.Get() => a 64-byte buffer
func (p *BufPool) Get() []byte {
	panic("not implemented")
}
