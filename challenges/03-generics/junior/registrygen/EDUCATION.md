# Registry

## Intuition

Go maps have no "insert if absent" primitive, so the check is the whole point of the method. Returning the bool lets callers fail loudly on duplicates.

## Approach

1. `NewRegistry`: allocate the map.
2. `Register`: return `false` when the key exists; otherwise store and return `true`.
3. `Lookup`: comma-ok.
4. `Len`: map length.

## Solution

```go
func NewRegistry[K comparable, V any]() *Registry[K, V] {
	return &Registry[K, V]{items: make(map[K]V)}
}

func (r *Registry[K, V]) Register(k K, v V) bool {
	if _, ok := r.items[k]; ok {
		return false
	}
	r.items[k] = v
	return true
}

func (r *Registry[K, V]) Lookup(k K) (V, bool) {
	v, ok := r.items[k]
	return v, ok
}

func (r *Registry[K, V]) Len() int {
	return len(r.items)
}
```

## Walkthrough

`Register(a, 2)` after `Register(a, 1)` finds the key present, returns `false`, and leaves `1` in place.

## Pitfalls

- Assigning first and checking afterwards, which has already overwritten.
- Returning `true` unconditionally.
- Forgetting to allocate the map, panicking on the first registration.
