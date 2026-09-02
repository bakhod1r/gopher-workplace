// Package ringcapbug — Gopher Workplace challenge.
package ringcapbug

// Ring is a fixed-capacity circular buffer.
type Ring[T any] struct {
	buf  []T
	head int
	n    int
}

// NewRing returns an empty ring holding at most capacity elements.
func NewRing[T any](capacity int) *Ring[T] {
	if capacity < 1 {
		capacity = 1
	}
	return &Ring[T]{buf: make([]T, capacity)}
}

// Slice returns the buffered elements, oldest first.
//
// The result is independent of the ring: appending to it must never
// write into the ring's storage, so its capacity equals its length.
//
// Examples:
//
//	Push(1); Push(2); Slice() => []T{1, 2}
func (r *Ring[T]) Slice() []T {
	// CHANGE CODE BELOW THIS LINE
	if r.n == 0 {
		return []T{}
	}
	end := r.head + r.n
	if end <= len(r.buf) {
		return r.buf[r.head:end]
	}
	out := make([]T, 0, r.n)
	out = append(out, r.buf[r.head:]...)
	out = append(out, r.buf[:end-len(r.buf)]...)
	return out
	// CHANGE CODE ABOVE THIS LINE
}

// Push adds v, overwriting the oldest element when the ring is full.
func (r *Ring[T]) Push(v T) {
	if r.n == len(r.buf) {
		r.buf[r.head] = v
		r.head = (r.head + 1) % len(r.buf)
		return
	}
	r.buf[(r.head+r.n)%len(r.buf)] = v
	r.n++
}

// Len reports how many elements the ring currently holds.
func (r *Ring[T]) Len() int {
	return r.n
}
