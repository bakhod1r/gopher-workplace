// Package mapmethodlesson — Gopher Workplace challenge.
package mapmethodlesson

// Slice is a named slice type of T that carries methods.
type Slice[T any] []T

// MapSlice applies f to every element and returns the results.
// It is a function, not a method, because the result element type
// is a second type parameter.
func MapSlice[T, U any](s Slice[T], f func(T) U) Slice[U] {
	// TODO(candidate): apply f to each element, collecting the results.
	panic("not implemented")
}

// Each calls f for every element, in order.
// Each can be a method because it introduces no new type.
func (s Slice[T]) Each(f func(T)) {
	// TODO(candidate): call f for each element in order.
	panic("not implemented")
}
