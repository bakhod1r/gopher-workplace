// Package ringbuffer — Gopher Workplace challenge.
package ringbuffer

import "errors"

// RingBuffer is a fixed-size circular buffer.
type RingBuffer struct {
	data []int
	head int
	tail int
	size int
}

// New creates a RingBuffer.
func New(cap int) *RingBuffer {
	return &RingBuffer{data: make([]int, cap)}
}

// Push adds an element. If full, returns error.
func (r *RingBuffer) Push(val int) error {
	// TODO(candidate): check if full (size == len(data)).
	// Write to tail, advance tail, increment size.
	_ = errors.New
	panic("not implemented")
}
