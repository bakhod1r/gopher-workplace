// Package ringbuffer — Gopher Workplace challenge.
package ringbuffer

import "errors"

// ErrFull is returned by Push when the buffer has no free slot.
var ErrFull = errors.New("ring buffer full")

// ErrEmpty is returned by Pop when the buffer holds no elements.
var ErrEmpty = errors.New("ring buffer empty")

// RingBuffer is a fixed-size circular buffer.
type RingBuffer struct {
	data []int
	head int
	tail int
	size int
}

// New creates a RingBuffer with room for cap elements.
func New(cap int) *RingBuffer {
	return &RingBuffer{data: make([]int, cap)}
}

// Len returns the number of buffered elements.
func (r *RingBuffer) Len() int { return r.size }

// Push appends val. It returns ErrFull, and changes nothing, when the buffer
// is full.
func (r *RingBuffer) Push(val int) error {
	// TODO(candidate): reject when size == len(r.data); otherwise write at
	// tail, advance tail modulo the capacity, and grow size.
	panic("not implemented")
}

// Pop removes and returns the oldest element. It returns ErrEmpty when the
// buffer holds nothing.
func (r *RingBuffer) Pop() (int, error) {
	// TODO(candidate): reject when size == 0; otherwise read at head, advance
	// head modulo the capacity, and shrink size.
	panic("not implemented")
}
