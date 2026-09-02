# Multi Map

## Intuition

A nil slice is a valid `append` target, so the absent-key case needs no branch. The copy in `Get` is what keeps the internal slice private.

## Approach

1. `NewMultiMap`: allocate the map.
2. `Add`: append and store the result back.
3. `Get`: copy the stored slice into a fresh one.

## Solution

```go
func NewMultiMap[K comparable, V any]() *MultiMap[K, V] {
	return &MultiMap[K, V]{items: make(map[K][]V)}
}

func (m *MultiMap[K, V]) Add(k K, v V) {
	m.items[k] = append(m.items[k], v)
}

func (m *MultiMap[K, V]) Get(k K) []V {
	out := make([]V, len(m.items[k]))
	copy(out, m.items[k])
	return out
}
```

## Walkthrough

`Add("a", 1)` reads a nil slice, appends to get `[1]`, and stores that back under `"a"`.

## Pitfalls

- Forgetting to assign the `append` result back into the map.
- Returning `m.items[k]` directly, which aliases the map's storage.
- Returning `nil` for an unknown key when an empty slice is expected.
