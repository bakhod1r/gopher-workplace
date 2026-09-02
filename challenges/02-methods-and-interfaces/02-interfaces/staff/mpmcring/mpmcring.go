// Package mpmcring — Gopher Workplace challenge.
package mpmcring

import "sync/atomic"

// Queue is a non-blocking queue.
type Queue interface {
	Enqueue(v int) bool
	Dequeue() (int, bool)
}

// slot is one ring cell with its sequence number.
type slot struct {
	seq atomic.Uint64
	val int
}

// Ring is a bounded multi-producer multi-consumer queue.
type Ring struct {
	mask uint64
	buf  []slot

	head atomic.Uint64 // next dequeue position
	tail atomic.Uint64 // next enqueue position
}

// NewRing returns a ring with the given capacity, rounded up to a power of two.
func NewRing(capacity int) *Ring {
	n := uint64(1)
	for n < uint64(capacity) {
		n <<= 1
	}
	r := &Ring{mask: n - 1, buf: make([]slot, n)}
	for i := range r.buf {
		r.buf[i].seq.Store(uint64(i))
	}
	return r
}

// Enqueue adds a value, returning false when the ring is full.
//
// Examples:
//
//	capacity 2; Enqueue, Enqueue, Enqueue => true, true, false
func (r *Ring) Enqueue(v int) bool {
	// TODO(candidate): claim a tail slot when seq == pos, write, publish pos+1.
	panic("not implemented")
}

// Dequeue removes a value, returning false when the ring is empty.
func (r *Ring) Dequeue() (int, bool) {
	// TODO(candidate): claim a head slot when seq == pos+1, read, publish pos+n.
	panic("not implemented")
}

// Cap returns the ring capacity.
func (r *Ring) Cap() int { return len(r.buf) }
