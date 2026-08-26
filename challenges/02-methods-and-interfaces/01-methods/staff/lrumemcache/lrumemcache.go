// Package lrumemcache — Gopher Workplace challenge.
package lrumemcache

type node struct {
	key        string
	val        int
	prev, next *node
}

// LRU cache.
type LRU struct {
	cap        int
	cache      map[string]*node
	head, tail *node
}

func New(cap int) *LRU {
	l := &LRU{cap: cap, cache: make(map[string]*node)}
	l.head = &node{}
	l.tail = &node{}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (l *LRU) remove(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

func (l *LRU) insert(n *node) {
	n.next = l.head.next
	n.prev = l.head
	l.head.next.prev = n
	l.head.next = n
}

// Get moves to front and returns value.
func (l *LRU) Get(key string) (int, bool) {
	// TODO(candidate): if found, remove and insert at head, return val, true
	panic("not implemented")
}

// Put inserts or updates, moving to front. Evicts tail if needed.
func (l *LRU) Put(key string, val int) {
	// TODO(candidate): if found, update val, remove and insert.
	// Else, create node, insert, add to map. If len(map) > cap, remove tail.prev, delete from map.
	panic("not implemented")
}
