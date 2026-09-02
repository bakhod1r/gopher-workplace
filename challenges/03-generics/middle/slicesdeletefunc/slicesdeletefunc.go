// Package slicesdeletefunc — Gopher Workplace challenge.
package slicesdeletefunc

// Purge returns s without the elements drop accepts,
// leaving the input untouched.
func Purge[T any](s []T, drop func(T) bool) []T {
	// TODO(candidate): clone, then delete the matching elements.
	panic("not implemented")
}
