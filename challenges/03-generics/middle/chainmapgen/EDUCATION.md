# Layered Configuration

## Intuition

The two methods traverse in opposite directions for the same reason: last-write-wins in a merge is the mirror image of first-hit-wins in a search.

## Approach

1. `Get`: return the first comma-ok hit.
2. `Flatten`: copy layers in reverse order into a fresh map.

## Solution

```go
func NewChain[K comparable, V any](layers ...map[K]V) *Chain[K, V] {
	return &Chain[K, V]{layers: layers}
}

func (c *Chain[K, V]) Get(k K) (V, bool) {
	for _, m := range c.layers {
		if v, ok := m[k]; ok {
			return v, true
		}
	}
	var zero V
	return zero, false
}

func (c *Chain[K, V]) Flatten() map[K]V {
	out := make(map[K]V)
	for i := len(c.layers) - 1; i >= 0; i-- {
		for k, v := range c.layers[i] {
			out[k] = v
		}
	}
	return out
}
```

## Walkthrough

Flattening `{a:1}` over `{a:2, b:3}` writes the low layer first, then overwrites `a` with `1`.

## Pitfalls

- Merging forward, so the lowest-priority layer wins.
- Testing values against the zero value instead of using `ok`.
- Returning one of the layers instead of a new map.
