// Package nilslicepitfall — Gopher Workplace challenge.
package nilslicepitfall

// Collect returns the accepted elements. It never returns nil.
func Collect[T any](s []T, keep func(T) bool) []T {
	// TODO(candidate): collect the accepted elements into a non-nil slice.
	panic("not implemented")
}

// IsNil reports whether s is nil, as opposed to merely empty.
func IsNil[T any](s []T) bool {
	// TODO(candidate): report whether the slice is nil.
	panic("not implemented")
}
