// Package ringstartbug — Gopher Workplace challenge.
package ringstartbug

// Ring is a fixed-capacity circular buffer of T.
// Use NewRing to create one.
type Ring[T any] struct {
	buf  []T
	head int
	n    int
}

// NewRing returns a ring holding at most size elements.
func NewRing[T any](size int) *Ring[T] {
	if size < 0 {
		size = 0
	}
	return &Ring[T]{buf: make([]T, size)}
}

// Add stores v, overwriting the oldest element when full.
// It is provided for you.
func (r *Ring[T]) Add(v T) {
	if len(r.buf) == 0 {
		return
	}
	r.buf[r.head] = v
	r.head = (r.head + 1) % len(r.buf)
	if r.n < len(r.buf) {
		r.n++
	}
}

// Items returns the buffered elements, oldest first.
func (r *Ring[T]) Items() []T {
	// CHANGE CODE BELOW THIS LINE
	out := make([]T, 0, r.n)
	if r.n == 0 {
		return out
	}
	start := (r.head - r.n) % len(r.buf)
	for i := 0; i < r.n; i++ {
		out = append(out, r.buf[(start+i)%len(r.buf)])
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
