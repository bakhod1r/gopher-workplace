// Package boxgen — Gopher Workplace challenge.
package boxgen

// Box holds at most one value of T.
// Its zero value is an empty box.
type Box[T any] struct {
	value  T
	filled bool
}

// Set stores v in the box.
func (b *Box[T]) Set(v T) {
	// TODO(candidate): store the value.
	panic("not implemented")
}

// Get returns the stored value and true, or the zero value and
// false when nothing has been stored.
func (b *Box[T]) Get() (T, bool) {
	// TODO(candidate): report the stored value, if any.
	panic("not implemented")
}
