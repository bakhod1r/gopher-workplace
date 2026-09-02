// Package preallocgen — Gopher Workplace challenge.
package preallocgen

// Build returns n elements produced by f, allocated in one go.
func Build[T any](n int, f func(int) T) []T {
	// TODO(candidate): allocate the exact capacity, then fill it.
	panic("not implemented")
}
