// Package derefgen — Gopher Workplace challenge.
package derefgen

// Deref returns *p, or def when p is nil.
//
// Examples:
//
//	Deref(Ptr(7), 0) => 7
//	Deref(nil, 0)    => 0
func Deref[T any](p *T, def T) T {
	// TODO(candidate): guard the nil pointer, then dereference.
	panic("not implemented")
}

// Ptr returns a pointer to a copy of v. It is provided for the tests.
func Ptr[T any](v T) *T {
	return &v
}
