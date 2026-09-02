// Package slicetypegen — Gopher Workplace challenge.
package slicetypegen

// Slice is a named slice type of T that carries methods.
type Slice[T any] []T

// Filter returns the elements keep accepts, as a Slice.
func (s Slice[T]) Filter(keep func(T) bool) Slice[T] {
	// TODO(candidate): collect the elements keep accepts.
	panic("not implemented")
}

// Len returns the number of elements.
func (s Slice[T]) Len() int {
	// TODO(candidate): report the length.
	panic("not implemented")
}
