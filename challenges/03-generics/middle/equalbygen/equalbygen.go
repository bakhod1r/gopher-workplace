// Package equalbygen — Gopher Workplace challenge.
package equalbygen

// EqualBy reports whether a and b have the same length and
// every matching pair satisfies eq.
func EqualBy[T, U any](a []T, b []U, eq func(T, U) bool) bool {
	// TODO(candidate): compare lengths, then each pair with eq.
	panic("not implemented")
}
