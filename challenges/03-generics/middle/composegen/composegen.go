// Package composegen — Gopher Workplace challenge.
package composegen

// Compose returns a function applying f then g.
func Compose[A, B, C any](f func(A) B, g func(B) C) func(A) C {
	// TODO(candidate): return a closure applying f then g.
	panic("not implemented")
}
