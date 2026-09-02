// Package currygen — Gopher Workplace challenge.
package currygen

// Curry2 turns a two-argument function into a chain of
// single-argument functions.
func Curry2[A, B, C any](f func(A, B) C) func(A) func(B) C {
	// TODO(candidate): return nested closures capturing each argument.
	panic("not implemented")
}
