// Package triegen — Gopher Workplace challenge.
package triegen

// Trie maps string keys to values of V, one rune per level.
// Use NewTrie to create one.
type Trie[V any] struct {
	children map[rune]*Trie[V]
	value    V
	set      bool
}

// NewTrie returns an empty trie.
func NewTrie[V any]() *Trie[V] {
	// TODO(candidate): allocate the root node.
	panic("not implemented")
}

// Insert stores v under key.
func (t *Trie[V]) Insert(key string, v V) {
	// TODO(candidate): walk the runes, creating nodes, then store the value.
	panic("not implemented")
}

// Get returns the value stored under key.
func (t *Trie[V]) Get(key string) (V, bool) {
	// TODO(candidate): walk the runes, then report the node's value.
	panic("not implemented")
}

// HasPrefix reports whether any stored key starts with prefix.
func (t *Trie[V]) HasPrefix(prefix string) bool {
	// TODO(candidate): walk the runes and report whether the path exists.
	panic("not implemented")
}
