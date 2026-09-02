// Package setmethodgen — Gopher Workplace challenge.
package setmethodgen

// Set is a collection of distinct values of T.
// Use NewSet to create one.
type Set[T comparable] struct {
	items map[T]struct{}
}

// NewSet returns a set holding the given values.
func NewSet[T comparable](vs ...T) *Set[T] {
	// TODO(candidate): allocate and fill the set.
	panic("not implemented")
}

// Union returns a new set holding the elements of s and other.
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	// TODO(candidate): return a new set with both sides' elements.
	panic("not implemented")
}

// Intersect returns a new set holding the shared elements.
func (s *Set[T]) Intersect(other *Set[T]) *Set[T] {
	// TODO(candidate): return a new set with the shared elements.
	panic("not implemented")
}

// Len returns the number of elements.
func (s *Set[T]) Len() int {
	// TODO(candidate): report the size.
	panic("not implemented")
}
