// Package dequegen — Gopher Workplace challenge.
package dequegen

// Deque supports adding at both ends and removing from the front.
// Its zero value is an empty deque.
type Deque[T any] struct {
	items []T
}

// PushFront adds v to the front.
func (d *Deque[T]) PushFront(v T) {
	// TODO(candidate): insert v at the front.
	panic("not implemented")
}

// PushBack adds v to the back.
func (d *Deque[T]) PushBack(v T) {
	// TODO(candidate): append v at the back.
	panic("not implemented")
}

// PopFront removes and returns the front element and true.
// It returns the zero value and false when empty.
func (d *Deque[T]) PopFront() (T, bool) {
	// TODO(candidate): remove and return the front element.
	panic("not implemented")
}
