// Package methodtypeparamgen — Gopher Workplace challenge.
package methodtypeparamgen

// Bag holds values of T with value semantics.
// Its zero value is an empty bag.
type Bag[T any] struct {
	items []T
}

// MapBag converts a bag of T into a bag of U.
// It is a function because U is a new type parameter.
func MapBag[T, U any](b Bag[T], f func(T) U) Bag[U] {
	// TODO(candidate): convert every item, returning a new bag.
	panic("not implemented")
}

// Add stores v in the bag and returns the bag for chaining.
// It can be a method because it introduces no new type.
func (b Bag[T]) Add(v T) Bag[T] {
	// TODO(candidate): return a bag with v appended.
	panic("not implemented")
}

// Items returns the stored values.
func (b Bag[T]) Items() []T {
	// TODO(candidate): return a copy of the stored values.
	panic("not implemented")
}
