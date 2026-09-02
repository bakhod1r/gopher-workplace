// Package scangen — Gopher Workplace challenge.
package scangen

// Scan returns the running accumulator after each element.
func Scan[T, A any](s []T, init A, f func(A, T) A) []A {
	// TODO(candidate): record the accumulator after each step.
	panic("not implemented")
}
