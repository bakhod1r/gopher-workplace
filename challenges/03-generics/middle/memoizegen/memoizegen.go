// Package memoizegen — Gopher Workplace challenge.
package memoizegen

// Memoize returns a function caching f's result per argument.
// The returned function is not safe for concurrent use.
func Memoize[K comparable, V any](f func(K) V) func(K) V {
	// TODO(candidate): return a closure caching results per argument.
	panic("not implemented")
}
