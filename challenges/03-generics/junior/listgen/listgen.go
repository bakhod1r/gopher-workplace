// Package listgen — Gopher Workplace challenge.
package listgen

// List is a growable sequence of T.
// Its zero value is an empty list.
type List[T any] struct {
	items []T
}

// Append adds v to the end of the list.
func (l *List[T]) Append(v T) {
	// TODO(candidate): append v to the items.
	panic("not implemented")
}

// At returns the element at index i and true.
// It returns the zero value and false for an out-of-range index.
func (l *List[T]) At(i int) (T, bool) {
	// TODO(candidate): bounds-check, then index.
	panic("not implemented")
}

// Len returns the number of elements.
func (l *List[T]) Len() int {
	// TODO(candidate): report how many items are stored.
	panic("not implemented")
}
