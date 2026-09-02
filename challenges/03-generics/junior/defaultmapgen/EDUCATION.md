# Default Map

## Intuition

Moving the fallback into the type means every caller of `Get` agrees on it automatically, which is the practical difference between the function and the type version.

## Approach

1. `NewDefaultMap`: allocate the map and store `def`.
2. `Put`: assign.
3. `Get`: comma-ok lookup, falling back to `m.def`.

## Solution

```go
func NewDefaultMap[K comparable, V any](def V) *DefaultMap[K, V] {
	return &DefaultMap[K, V]{items: make(map[K]V), def: def}
}

func (m *DefaultMap[K, V]) Put(k K, v V) {
	m.items[k] = v
}

func (m *DefaultMap[K, V]) Get(k K) V {
	if v, ok := m.items[k]; ok {
		return v
	}
	return m.def
}
```

## Walkthrough

`Put("x", 0); Get("x")` returns `0`: the key exists, so the fallback is not consulted.

## Pitfalls

- Returning `m.def` whenever the stored value is zero.
- Forgetting to allocate the map in the constructor.
- Storing the default as a pointer and mutating it later by accident.
