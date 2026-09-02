// Package aliasgen — Gopher Workplace challenge.
package aliasgen

// Index is a generic alias for a set of keys.
// Being an alias, Index[K] is the same type as map[K]struct{}.
type Index[K comparable] = map[K]struct{}

// NewIndex returns an empty Index.
func NewIndex[K comparable]() Index[K] {
	// TODO(candidate): return an allocated index.
	panic("not implemented")
}

// Mark records k in the index.
func Mark[K comparable](ix Index[K], k K) {
	// TODO(candidate): record the key.
	panic("not implemented")
}

// Marked reports whether k was recorded.
func Marked[K comparable](ix Index[K], k K) bool {
	// TODO(candidate): report whether the key is recorded.
	panic("not implemented")
}
