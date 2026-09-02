// Package slicesclipgen — Gopher Workplace challenge.
package slicesclipgen

// Freeze returns s with its spare capacity removed, so appending
// to the result cannot overwrite anything beyond its length.
func Freeze[T any](s []T) []T {
	// TODO(candidate): remove the spare capacity.
	panic("not implemented")
}
