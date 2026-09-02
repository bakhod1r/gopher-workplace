# The Cache That Walks Itself

## Intuition

The entry is already reachable in constant time through the map, and its neighbours are reachable through its own `prev`/`next` pointers. Rebuilding the whole list to move one node throws that away and makes every lookup cost `O(n)` — plus an allocation proportional to the cache size.

## Approach

1. Look the node up in the map.
2. Unlink it from its two neighbours.
3. Push it to the front of the list.

## Solution

```go
func (c *LRU[K, V]) Get(k K) (V, bool) {
	n, ok := c.m[k]
	if !ok {
		var zero V
		return zero, false
	}
	c.unlink(n)
	c.pushFront(n)
	return n.val, true
}

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

func (c *LRU[K, V]) Keys() []K {
	out := make([]K, 0, len(c.m))
	for p := c.head.next; p != c.tail; p = p.next {
		out = append(out, p.key)
	}
	return out
}

func (c *LRU[K, V]) Len() int {
	return len(c.m)
}
```

## Walkthrough

On a 16384-entry cache, 150000 lookups walk and re-link roughly 2.5 billion nodes and allocate 150000 slices of 16384 pointers; the constant-time version touches about 600000 pointers in total.

## Pitfalls

- Keeping recency in a slice and splicing it — same `O(n)` cost in different clothes.
- Skipping the unlink before the push, which corrupts the list into a cycle.
