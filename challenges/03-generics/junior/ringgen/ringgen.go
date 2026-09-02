// Package ringgen — Gopher Workplace challenge.
package ringgen

// Ring keeps at most size elements, dropping the oldest first.
// Use NewRing to create one.
type Ring[T any] struct {
	items []T
	size  int
}

// NewRing returns a ring holding at most size elements.
// It returns a ring of capacity 0 when size <= 0.
func NewRing[T any](size int) *Ring[T] {
	// TODO(candidate): build a ring with the given capacity.
	panic("not implemented")
}

// Add appends v, dropping the oldest element when the ring is full.
func (r *Ring[T]) Add(v T) {
	// TODO(candidate): append, then drop the oldest element if over capacity.
	panic("not implemented")
}

// Items returns the buffered elements, oldest first.
func (r *Ring[T]) Items() []T {
	// TODO(candidate): return a copy of the buffered elements.
	panic("not implemented")
}
