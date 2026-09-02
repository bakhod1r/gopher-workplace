# Bounded Cache

## Intuition

An unbounded cache is a memory leak with good intentions. The ceiling must be enforced on the write path, before insertion, so the invariant `Len() <= Max` never breaks even momentarily.

## Approach

1. Return the cached value on a hit.
2. On a miss, fetch from the inner source.
3. When `Max <= 0`, return without caching at all.
4. When `len(order) >= Max`, drop `order[0]` from both the slice and the map.
5. Insert the new key into the map and append it to `order`.

## Solution

```go
func (c *Cache) Get(key string) string {
	if v, ok := c.entries[key]; ok {
		return v
	}
	v := c.inner.Get(key)
	if c.Max <= 0 {
		return v
	}
	if len(c.order) >= c.Max {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.entries, oldest)
	}
	c.entries[key] = v
	c.order = append(c.order, key)
	return v
}

func (c *Cache) Len() int { return len(c.entries) }
```

## Walkthrough

With `Max` 2 and keys a, b, c: inserting `c` sees `len(order) == 2`, evicts `a` from both structures, then stores `c`. `Len` never reads 3.

## Pitfalls

- Inserting first and evicting afterwards — the ceiling is briefly violated, which matters under memory pressure.
- Deleting from the map but not from `order`, so the slice grows unbounded and becomes the leak.
- Forgetting `Max <= 0`, which makes `order[0]` panic on an empty slice.
