// Package lruo1bug — Gopher Workplace challenge.
package lruo1bug

type lnode[K comparable, V any] struct {
	key        K
	val        V
	prev, next *lnode[K, V]
}

// LRU is a fixed-capacity least-recently-used cache.
type LRU[K comparable, V any] struct {
	cap  int
	m    map[K]*lnode[K, V]
	head *lnode[K, V]
	tail *lnode[K, V]
}

// NewLRU returns an empty cache holding at most capacity entries.
func NewLRU[K comparable, V any](capacity int) *LRU[K, V] {
	if capacity < 1 {
		capacity = 1
	}
	h := &lnode[K, V]{}
	t := &lnode[K, V]{}
	h.next = t
	t.prev = h
	return &LRU[K, V]{cap: capacity, m: make(map[K]*lnode[K, V]), head: h, tail: t}
}

func (c *LRU[K, V]) unlink(n *lnode[K, V]) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (c *LRU[K, V]) pushFront(n *lnode[K, V]) {
	n.prev = c.head
	n.next = c.head.next
	c.head.next.prev = n
	c.head.next = n
}

// Get returns the value for k and true, promoting the entry to most
// recently used.
//
// It must run in constant time regardless of the cache size.
//
// Examples:
//
//	Put(1, "a"); Get(1) => "a", true
func (c *LRU[K, V]) Get(k K) (V, bool) {
	// CHANGE CODE BELOW THIS LINE
	n, ok := c.m[k]
	if !ok {
		var zero V
		return zero, false
	}
	rest := make([]*lnode[K, V], 0, len(c.m))
	for p := c.head.next; p != c.tail; p = p.next {
		if p != n {
			rest = append(rest, p)
		}
	}
	c.head.next = c.tail
	c.tail.prev = c.head
	for i := len(rest) - 1; i >= 0; i-- {
		c.pushFront(rest[i])
	}
	c.pushFront(n)
	return n.val, true
	// CHANGE CODE ABOVE THIS LINE
}

// Put stores v under k, evicting the least recently used entry when full.
func (c *LRU[K, V]) Put(k K, v V) {
	if n, ok := c.m[k]; ok {
		n.val = v
		c.unlink(n)
		c.pushFront(n)
		return
	}
	if len(c.m) == c.cap {
		old := c.tail.prev
		c.unlink(old)
		delete(c.m, old.key)
	}
	n := &lnode[K, V]{key: k, val: v}
	c.m[k] = n
	c.pushFront(n)
}

// Keys returns the cached keys, most recently used first.
func (c *LRU[K, V]) Keys() []K {
	out := make([]K, 0, len(c.m))
	for p := c.head.next; p != c.tail; p = p.next {
		out = append(out, p.key)
	}
	return out
}

// Len reports how many entries the cache holds.
func (c *LRU[K, V]) Len() int {
	return len(c.m)
}
