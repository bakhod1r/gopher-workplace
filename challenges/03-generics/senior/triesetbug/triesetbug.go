// Package triesetbug — Gopher Workplace challenge.
package triesetbug

type node[T comparable] struct {
	kids map[T]*node[T]
	end  bool
}

// Trie stores sequences of comparable elements.
type Trie[T comparable] struct {
	r *node[T]
}

func (t *Trie[T]) root() *node[T] {
	if t.r == nil {
		t.r = &node[T]{kids: make(map[T]*node[T])}
	}
	return t.r
}

// Insert adds the sequence to the trie.
func (t *Trie[T]) Insert(seq []T) {
	// CHANGE CODE BELOW THIS LINE
	n := t.root()
	n.end = true
	for _, e := range seq {
		next, ok := n.kids[e]
		if !ok {
			next = &node[T]{kids: make(map[T]*node[T])}
			n.kids[e] = next
		}
		n = next
	}
	// CHANGE CODE ABOVE THIS LINE
}

// Contains reports whether the exact sequence was inserted.
func (t *Trie[T]) Contains(seq []T) bool {
	n := t.root()
	for _, e := range seq {
		next, ok := n.kids[e]
		if !ok {
			return false
		}
		n = next
	}
	return n.end
}
