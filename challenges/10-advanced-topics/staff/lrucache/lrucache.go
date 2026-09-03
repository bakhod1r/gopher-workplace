// Package lrucache — Gopher Workplace challenge.
package lrucache

import (
	"container/list"
	"sync"
)

// entry is one cached pair, stored in the list.
type entry struct {
	key string
	val int
}

// LRU is a bounded, concurrency-safe least-recently-used cache.
type LRU struct {
	mu    sync.Mutex
	limit int
	ll    *list.List
	items map[string]*list.Element
}

// NewLRU returns a cache holding at most limit entries.
func NewLRU(limit int) *LRU {
	if limit < 1 {
		limit = 1
	}
	return &LRU{limit: limit, ll: list.New(), items: make(map[string]*list.Element, limit)}
}

// Put stores a value, evicting the least recently used entry if needed.
func (c *LRU) Put(key string, val int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if el, ok := c.items[key]; ok {
		c.ll.MoveToFront(el)
		el.Value.(*entry).val = val
		return
	}
	if c.ll.Len() >= c.limit {
		oldest := c.ll.Back()
		if oldest != nil {
			c.ll.Remove(oldest)
			delete(c.items, oldest.Value.(*entry).key)
		}
	}
	c.items[key] = c.ll.PushFront(&entry{key: key, val: val})
}

// Len reports how many entries the cache holds.
func (c *LRU) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.ll.Len()
}

// Get returns the value for key and marks it as the most recently used.
//
// A cache that evicts by insertion order throws away the hot entries; the
// whole point of LRU is that a hit moves the entry to the front.
//
// Examples:
//
//	c.Put("a", 1); c.Get("a") => 1, true and "a" becomes newest
func (c *LRU) Get(key string) (int, bool) {
	panic("not implemented")
}
