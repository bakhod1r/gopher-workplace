# LRU That Is Really FIFO

## Intuition

Eviction order is driven entirely by the `order` list. If reads never reorder it, the list records insertion order and the oldest inserted key is dropped no matter how often it is read.

## Approach

1. On a hit, promote the key before returning the value.
2. Leave misses alone — promoting an absent key would corrupt the order list.

## Solution

```go
func NewLRU[K comparable, V any](size int) *LRU[K, V] {
	if size < 0 {
		size = 0
	}
	return &LRU[K, V]{items: make(map[K]V), order: make([]K, 0), size: size}
}

func (c *LRU[K, V]) Get(k K) (V, bool) {
	v, ok := c.items[k]
	if !ok {
		var zero V
		return zero, false
	}
	c.touch(k)
	return v, true
}

func (c *LRU[K, V]) Put(k K, v V) {
	if c.size == 0 {
		return
	}
	c.items[k] = v
	c.touch(k)
	if len(c.order) > c.size {
		oldest := c.order[0]
		c.order = c.order[1:]
		delete(c.items, oldest)
	}
}

func (c *LRU[K, V]) touch(k K) {
	for i, key := range c.order {
		if key == k {
			c.order = append(c.order[:i], c.order[i+1:]...)
			break
		}
	}
	c.order = append(c.order, k)
}
```

## Walkthrough

Reading `a` in a full cache must move it behind `b`, so the next insert evicts `b`.

## Pitfalls

- Promoting on a miss as well, which inserts a key with no value.
- Promoting inside the map lookup and losing the `ok` result.
- Deciding the cache "works well enough" because tests only measure hits.
