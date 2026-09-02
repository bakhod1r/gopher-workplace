// Package multisetgen — Gopher Workplace challenge.
package multisetgen

// Bag is a multiset of T.
// Use NewBag to create one.
type Bag[T comparable] struct {
	counts map[T]int
}

// NewBag returns an empty multiset.
func NewBag[T comparable]() *Bag[T] {
	// TODO(candidate): allocate the count map.
	panic("not implemented")
}

// Add records one more occurrence of v.
func (b *Bag[T]) Add(v T) {
	// TODO(candidate): increase the count for v.
	panic("not implemented")
}

// Remove drops one occurrence of v, reporting whether one existed.
func (b *Bag[T]) Remove(v T) bool {
	// TODO(candidate): decrease the count, deleting the key at zero.
	panic("not implemented")
}

// Count returns how many occurrences of v are stored.
func (b *Bag[T]) Count(v T) int {
	// TODO(candidate): report the count for v.
	panic("not implemented")
}

// Distinct returns how many distinct values are stored.
func (b *Bag[T]) Distinct() int {
	// TODO(candidate): report the number of distinct values.
	panic("not implemented")
}
