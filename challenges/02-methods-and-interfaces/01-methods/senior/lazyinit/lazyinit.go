// Package lazyinit — Gopher Workplace challenge.
package lazyinit

// LazyString initializes its value on first access.
type LazyString struct {
	val  *string
	init func() string
}

// New creates a lazy string.
func New(init func() string) *LazyString {
	return &LazyString{init: init}
}

// String returns the value, calling init() only on the first call.
func (l *LazyString) String() string {
	// TODO(candidate): if val is nil, call init and store pointer to result.
	// Then return the dereferenced string.
	panic("not implemented")
}
