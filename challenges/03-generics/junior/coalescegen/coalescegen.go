// Package coalescegen — Gopher Workplace challenge.
package coalescegen

// Coalesce returns the first non-zero argument.
// It returns the zero value when every argument is zero.
//
// Examples:
//
//	Coalesce(0, 0, 5) => 5
//	Coalesce(0, 0)    => 0
func Coalesce[T comparable](vals ...T) T {
	// TODO(candidate): return the first argument that differs from the zero value.
	panic("not implemented")
}
