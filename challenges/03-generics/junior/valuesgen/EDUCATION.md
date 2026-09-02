# Map Values

## Intuition

Duplicate values are kept: unlike keys, map values are not unique, so the output length always equals `len(m)`.

## Approach

1. Allocate `out` as `[]V` with capacity `len(m)`.
2. Range over `m`, appending each value.
3. Return `out`.

## Solution

```go
func Values[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, v := range m {
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`Values(map[string]int{"a": 1, "b": 1})` returns two elements, both `1` — values are not deduplicated.

## Pitfalls

- Deduplicating the values, which changes the length.
- Allocating `[]K` instead of `[]V`.
- Relying on map order to match a fixed expected slice.
