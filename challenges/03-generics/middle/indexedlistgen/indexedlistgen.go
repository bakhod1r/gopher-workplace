// Package indexedlistgen — Gopher Workplace challenge.
package indexedlistgen

// Indexed is an ordered list of distinct values with O(1) membership.
// Use NewIndexed to create one.
type Indexed[T comparable] struct {
	items []T
	index map[T]int
}

// NewIndexed returns an empty list.
func NewIndexed[T comparable]() *Indexed[T] {
	// TODO(candidate): allocate the slice and the index.
	panic("not implemented")
}

// Append adds v unless it is already present, reporting whether
// it was added.
func (l *Indexed[T]) Append(v T) bool {
	// TODO(candidate): append only when v is new, keeping the index in step.
	panic("not implemented")
}

// Has reports whether v is present, in constant time.
func (l *Indexed[T]) Has(v T) bool {
	// TODO(candidate): answer from the index.
	panic("not implemented")
}

// At returns the element at position i.
func (l *Indexed[T]) At(i int) (T, bool) {
	// TODO(candidate): bounds-check, then index.
	panic("not implemented")
}
