# A Cache That Cannot Grow

## Intuition

The map answers lookups; the slice remembers arrival order. Eviction pops the front of the slice and deletes that key.

## Approach

1. Lazily create the map, and do nothing at all when `Cap <= 0`.
2. `Put` overwrites in place when the key exists; otherwise it evicts if full, then inserts and records the order.
3. `Get` is a map read plus a counter.

## Solution

```go
func (c *Cache) Get(key string) (int, bool) {
	v, ok := c.items[key]
	if ok {
		c.hits++
	} else {
		c.miss++
	}
	return v, ok
}

func (c *Cache) Put(key string, v int) {
	if c.Cap <= 0 {
		return
	}
	if _, ok := c.items[key]; ok {
		c.items[key] = v
		return
	}
	if c.items == nil {
		c.items = make(map[string]int, c.Cap)
	}
	if len(c.order) >= c.Cap {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.items, oldest)
	}
	c.items[key] = v
	c.order = append(c.order, key)
}

func (c *Cache) Len() int { return len(c.items) }

func (c *Cache) Stats() (int, int) { return c.hits, c.miss }
```

## Walkthrough

Reading a nil map is legal, so `Get` needs no initialisation check — only `Put` does. The overwrite branch returns before touching `order`, which is what keeps an update from behaving like a fresh insert.

## Pitfalls

- Appending to `order` on an overwrite, which grows the slice past `Cap` and evicts live entries.
- Evicting after inserting, which briefly exceeds the capacity and can evict the entry just added.
- Calling this an LRU; a hit here does not protect an entry from eviction.
