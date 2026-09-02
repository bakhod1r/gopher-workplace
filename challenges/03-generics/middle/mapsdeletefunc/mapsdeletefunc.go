// Package mapsdeletefunc — Gopher Workplace challenge.
package mapsdeletefunc

// Prune returns a copy of m without the entries drop accepts.
func Prune[K comparable, V any](m map[K]V, drop func(K, V) bool) map[K]V {
	// TODO(candidate): clone, then delete the matching entries.
	panic("not implemented")
}
