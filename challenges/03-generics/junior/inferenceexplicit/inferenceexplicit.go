// Package inferenceexplicit — Gopher Workplace challenge.
package inferenceexplicit

// Empty returns an empty slice of T.
// The caller must supply the type argument.
func Empty[T any]() []T {
	// TODO(candidate): return an empty slice of T.
	panic("not implemented")
}

// ZeroOf returns the zero value of T.
// The caller must supply the type argument.
func ZeroOf[T any]() T {
	// TODO(candidate): return the zero value of T.
	panic("not implemented")
}

// Wrap returns a one-element slice holding v.
// Here T is inferred from the argument.
func Wrap[T any](v T) []T {
	// TODO(candidate): return a slice holding just v.
	panic("not implemented")
}
