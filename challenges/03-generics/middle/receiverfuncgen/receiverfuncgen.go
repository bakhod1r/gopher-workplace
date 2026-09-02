// Package receiverfuncgen — Gopher Workplace challenge.
package receiverfuncgen

// Box holds a single value of T.
type Box[T any] struct {
	value T
}

// Update applies f to the boxed value in place.
func Update[T any](b *Box[T], f func(T) T) {
	// TODO(candidate): replace the boxed value with f applied to it.
	panic("not implemented")
}

// Convert returns a new box holding f applied to the value.
// It must be a function: U is a new type parameter.
func Convert[T, U any](b *Box[T], f func(T) U) *Box[U] {
	// TODO(candidate): return a new box holding the converted value.
	panic("not implemented")
}

// Get returns the boxed value.
func (b *Box[T]) Get() T {
	// TODO(candidate): return the stored value.
	panic("not implemented")
}
