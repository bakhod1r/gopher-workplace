// Package mustgen — Gopher Workplace challenge.
package mustgen

// Must returns v when ok is true, and panics otherwise.
// It is for package initialisation and tests, not request paths.
func Must[T any](v T, ok bool) T {
	// TODO(candidate): return v, or panic when ok is false.
	panic("not implemented")
}

// Lookup returns m[k] with a presence flag. It is provided
// so the tests have something to wrap.
func Lookup[K comparable, V any](m map[K]V, k K) (V, bool) {
	v, ok := m[k]
	return v, ok
}
