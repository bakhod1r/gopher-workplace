# A Cache That Cannot Outgrow Its Limit

## Intuition

Bounded memory needs two things a map does not give you: a rule for what to drop and ownership of what you keep. Both are decided at insert time, and only for keys that are actually new.

## Approach

1. Copy `val` into a private slice before taking the lock.
2. Under the lock, check whether the key exists.
3. For a new key at capacity, drop the front of the order slice and delete it from the map.
4. Record the new key in the order and store the copy.

## Solution

```go
import "sync"

// Cache is a bounded, concurrency-safe byte cache with FIFO eviction.
type Cache struct {
	mu    sync.Mutex
	limit int
	items map[string][]byte
	order []string
}

// NewCache returns a cache holding at most limit entries.
func NewCache(limit int) *Cache {
	if limit < 1 {
		limit = 1
	}
	return &Cache{limit: limit, items: make(map[string][]byte, limit), order: make([]string, 0, limit)}
}

// Get returns the stored bytes for key, if present.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	v, ok := c.items[key]
	return v, ok
}

// Len reports how many entries the cache holds.
func (c *Cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return len(c.items)
}

// Put stores a copy of val under key, evicting the oldest entry when the
// cache is at capacity.
//
// The stored value must own its bytes — callers reuse their buffers — and
// the cache must never hold more than limit entries.
//
// Examples:
//
// 	c := NewCache(2); c.Put("a", v) => Get("a") returns a copy of v
func (c *Cache) Put(key string, val []byte) {
	owned := make([]byte, len(val))
	copy(owned, val)

	c.mu.Lock()
	defer c.mu.Unlock()

	if _, exists := c.items[key]; !exists {
		if len(c.order) >= c.limit {
			oldest := c.order[0]
			c.order = append(c.order[:0], c.order[1:]...)
			delete(c.items, oldest)
		}
		c.order = append(c.order, key)
	}
	c.items[key] = owned
}
```

## Walkthrough

With a limit of 2 and puts a, b, c: c is new and the cache is full, so a is evicted. Putting a again after that would evict b — but putting b again would only overwrite.

## Pitfalls

- Appending to `order` on an overwrite, which lets the same key sit in the order twice.
- Storing `val` directly, which is invisible until the caller reuses the buffer.
- Holding the lock across the copy, which serialises every writer on a memcpy.
