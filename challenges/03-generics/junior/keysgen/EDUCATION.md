# Map Keys

## Intuition

The key type parameter must be `comparable` because Go requires map keys to support `==`. The value type stays `any` since this function never inspects a value.

## Approach

1. Allocate `out` with capacity `len(m)`.
2. Range over `m`, appending each key.
3. Return `out`.

## Solution

```go
func Keys[K comparable, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	return out
}
```

## Walkthrough

`Keys(map[int]bool{1: true, 2: false})` may return `[1 2]` or `[2 1]`; both are correct, so the test sorts first.

## Pitfalls

- Declaring `[K any, V any]` — the map type itself will not compile.
- Assuming a stable order and writing tests (or code) that depend on it.
- Using `for k, v := range m` and leaving `v` unused, which fails to compile.
