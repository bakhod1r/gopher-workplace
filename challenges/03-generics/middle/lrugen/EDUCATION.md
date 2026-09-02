# LRU Cache

## Intuition

Promotion on read is what makes the policy "least recently *used*". Without it, a hot key inserted long ago would still be evicted first.

## Approach

1. `NewLRU`: clamp the size and allocate.
2. `Get`: return `false` on a miss; otherwise promote and return the value.
3. `Put`: store, promote, and evict the front key while over capacity.

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

With size 2 holding `a, b`, reading `a` moves it behind `b`, so inserting `c` evicts `b`.

## Pitfalls

- Skipping the promotion in `Get`, which silently degrades to FIFO.
- Promoting on a miss, which inserts a key with no value.
- Evicting before inserting, which can drop the entry you just stored.
