# Bounded Cache

## Intuition

Keeping the order list free of duplicates is what makes the eviction count match the map size. Without that check, updating one key repeatedly would evict live entries.

## Approach

1. `NewCache`: clamp the size, allocate both structures.
2. `Put`: record the key in `order` only when new, store the value, then evict the front key while over capacity.
3. `Get`: comma-ok lookup.

## Solution

```go
func NewCache[K comparable, V any](size int) *Cache[K, V] {
	if size < 0 {
		size = 0
	}
	return &Cache[K, V]{items: make(map[K]V), order: make([]K, 0), size: size}
}

func (c *Cache[K, V]) Put(k K, v V) {
	if c.size == 0 {
		return
	}
	if _, ok := c.items[k]; !ok {
		c.order = append(c.order, k)
	}
	c.items[k] = v
	if len(c.order) > c.size {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.items, oldest)
	}
}

func (c *Cache[K, V]) Get(k K) (V, bool) {
	v, ok := c.items[k]
	return v, ok
}
```

## Walkthrough

With size 2, `Put(c,3)` makes `order` `[a b c]`, so `a` is evicted and `Get(a)` reports `false`.

## Pitfalls

- Appending to `order` on every `Put`, so updates evict live keys.
- Deleting from the map without dropping the key from `order`.
- Panicking on a zero-size cache instead of storing nothing.
