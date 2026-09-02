// Package ringidxgen — Gopher Workplace challenge.
package ringidxgen

// Ring is a fixed-capacity circular buffer of T.
// Use NewRing to create one.
type Ring[T any] struct {
	buf  []T
	head int
	n    int
}

// NewRing returns a ring holding at most size elements.
func NewRing[T any](size int) *Ring[T] {
	// TODO(candidate): allocate fixed storage.
	panic("not implemented")
}

// Add stores v, overwriting the oldest element when full.
func (r *Ring[T]) Add(v T) {
	// TODO(candidate): write at the head, advancing it modulo the capacity.
	panic("not implemented")
}

// Items returns the buffered elements, oldest first.
func (r *Ring[T]) Items() []T {
	// TODO(candidate): read n elements ending at the head, oldest first.
	panic("not implemented")
}
