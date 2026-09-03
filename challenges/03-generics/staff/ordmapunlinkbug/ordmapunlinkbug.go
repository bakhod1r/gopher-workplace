// Package ordmapunlinkbug — Gopher Workplace challenge.
package ordmapunlinkbug

type onode[K comparable, V any] struct {
	key        K
	val        V
	prev, next *onode[K, V]
}

// OrderedMap is a map that remembers insertion order.
type OrderedMap[K comparable, V any] struct {
	idx        map[K]*onode[K, V]
	head, tail *onode[K, V]
}

// Delete removes k and reports whether it was present.
// The remaining entries keep their insertion order in both directions.
func (m *OrderedMap[K, V]) Delete(k K) bool {
	// CHANGE CODE BELOW THIS LINE
	n, ok := m.idx[k]
	if !ok {
		return false
	}
	delete(m.idx, k)
	if n.prev != nil {
		n.prev.next = n.next
	} else {
		m.head = n.next
	}
	return true
	// CHANGE CODE ABOVE THIS LINE
}

// Set stores v under k, appending k at the end on first insert.
func (m *OrderedMap[K, V]) Set(k K, v V) {
	if m.idx == nil {
		m.idx = make(map[K]*onode[K, V])
	}
	if n, ok := m.idx[k]; ok {
		n.val = v
		return
	}
	n := &onode[K, V]{key: k, val: v, prev: m.tail}
	if m.tail != nil {
		m.tail.next = n
	} else {
		m.head = n
	}
	m.tail = n
	m.idx[k] = n
}

// Get returns the value stored under k.
func (m *OrderedMap[K, V]) Get(k K) (V, bool) {
	if n, ok := m.idx[k]; ok {
		return n.val, true
	}
	var zero V
	return zero, false
}

// Len reports how many entries the map holds.
func (m *OrderedMap[K, V]) Len() int {
	return len(m.idx)
}

// Keys lists the keys in insertion order.
func (m *OrderedMap[K, V]) Keys() []K {
	out := make([]K, 0, len(m.idx))
	for n := m.head; n != nil; n = n.next {
		out = append(out, n.key)
	}
	return out
}

// RevKeys lists the keys in reverse insertion order.
func (m *OrderedMap[K, V]) RevKeys() []K {
	out := make([]K, 0, len(m.idx))
	for n := m.tail; n != nil; n = n.prev {
		out = append(out, n.key)
	}
	return out
}
