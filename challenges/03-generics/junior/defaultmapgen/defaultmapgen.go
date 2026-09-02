// Package defaultmapgen — Gopher Workplace challenge.
package defaultmapgen

// DefaultMap is a map that returns a fallback value for unknown keys.
// Use NewDefaultMap to create one.
type DefaultMap[K comparable, V any] struct {
	items map[K]V
	def   V
}

// NewDefaultMap returns a map that yields def for unknown keys.
func NewDefaultMap[K comparable, V any](def V) *DefaultMap[K, V] {
	// TODO(candidate): store the fallback and allocate the map.
	panic("not implemented")
}

// Put stores v under k.
func (m *DefaultMap[K, V]) Put(k K, v V) {
	// TODO(candidate): store the value.
	panic("not implemented")
}

// Get returns the stored value, or the fallback for an unknown key.
func (m *DefaultMap[K, V]) Get(k K) V {
	// TODO(candidate): look up k, falling back to the stored default.
	panic("not implemented")
}
