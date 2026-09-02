// Package optionalgen — Gopher Workplace challenge.
package optionalgen

// Optional holds a value of T, or nothing.
// Its zero value is empty.
type Optional[T any] struct {
	value   T
	present bool
}

// Some returns an Optional holding v.
func Some[T any](v T) Optional[T] {
	// TODO(candidate): build a present Optional.
	panic("not implemented")
}

// None returns an empty Optional.
func None[T any]() Optional[T] {
	// TODO(candidate): build an absent Optional.
	panic("not implemented")
}

// Or returns the held value, or def when the Optional is empty.
func (o Optional[T]) Or(def T) T {
	// TODO(candidate): return the value or the fallback.
	panic("not implemented")
}
