// Package io_uring — Gopher Workplace challenge.
package io_uring

// ringSize is the fixed capacity of the submission queue.
const ringSize = 4

// Ring is a fixed-capacity submission queue backed by a circular buffer.
// head is the next slot to read, tail the next slot to write, and count the
// number of entries currently queued.
type Ring struct {
	buf   [ringSize]int
	head  int
	tail  int
	count int
}

// Submit enqueues op. It reports false and drops op when the ring is full.
func (r *Ring) Submit(op int) bool {
	// TODO(candidate): reject when full, else write at tail, advance tail
	// modulo ringSize, and grow count.
	panic("not implemented")
}

// Complete dequeues the oldest entry. The second result is false when the
// ring is empty.
func (r *Ring) Complete() (int, bool) {
	// TODO(candidate): reject when empty, else read at head, advance head
	// modulo ringSize, and shrink count.
	panic("not implemented")
}

// Len returns the number of queued entries.
func (r *Ring) Len() int { return r.count }
