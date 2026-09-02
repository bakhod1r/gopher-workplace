// Package lrucache — Gopher Workplace challenge.
package lrucache

// node is one entry in the recency list.
type node struct {
	key   string
	value string
	prev  *node
	next  *node
}

// LRU is a fixed-capacity least-recently-used cache.
type LRU struct {
	Cap   int
	items map[string]*node
	head  *node // sentinel: head.next is the most recent
	tail  *node // sentinel: tail.prev is the least recent
}

// NewLRU returns an empty cache with the given capacity.
func NewLRU(capacity int) *LRU {
	head, tail := &node{}, &node{}
	head.next, tail.prev = tail, head
	return &LRU{
		Cap:   capacity,
		items: make(map[string]*node, capacity),
		head:  head,
		tail:  tail,
	}
}

// unlink removes n from the recency list.
func (l *LRU) unlink(n *node) {
	n.prev.next = n.next
	n.next.prev = n.prev
}

// pushFront moves n to the most-recent position.
func (l *LRU) pushFront(n *node) {
	n.prev = l.head
	n.next = l.head.next
	l.head.next.prev = n
	l.head.next = n
}

// Get returns the value and marks the entry as most recently used.
//
// Examples:
//
//	Get on a hit  => value, true
//	Get on a miss => "", false
func (l *LRU) Get(key string) (string, bool) {
	// TODO(candidate): O(1) lookup, then touch.
	panic("not implemented")
}

// Put stores a value, evicting the least recently used entry when full.
func (l *LRU) Put(key, value string) {
	// TODO(candidate): update-and-touch, or insert with eviction.
	panic("not implemented")
}

// Len returns how many entries are cached.
func (l *LRU) Len() int {
	// TODO(candidate): entry count.
	panic("not implemented")
}
