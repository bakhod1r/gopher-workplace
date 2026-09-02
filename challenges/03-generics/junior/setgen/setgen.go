// Package setgen — Gopher Workplace challenge.
package setgen

// Set is a collection of distinct values of T.
// Use NewSet to create one.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet returns an empty, ready-to-use set.
func NewSet[T comparable]() *Set[T] {
	// TODO(candidate): return a set with its map allocated.
	panic("not implemented")
}

// Add stores v in the set. Adding twice has no extra effect.
func (s *Set[T]) Add(v T) {
	// TODO(candidate): store v as a key.
	panic("not implemented")
}

// Has reports whether v is in the set.
func (s *Set[T]) Has(v T) bool {
	// TODO(candidate): report whether the key exists.
	panic("not implemented")
}

// Len returns the number of distinct elements.
func (s *Set[T]) Len() int {
	// TODO(candidate): report how many keys are stored.
	panic("not implemented")
}
