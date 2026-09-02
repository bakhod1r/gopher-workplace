# Fallback Chain

## Intuition

This is the map version of `Coalesce`, and it fixes `Coalesce`'s weakness: presence is tracked by the map, so a deliberately stored zero shadows lower layers correctly.

## Approach

1. Scan the maps in order.
2. Return the value on the first comma-ok hit.
3. Return the zero value and `false` after the loop.

## Solution

```go
func Lookup[K comparable, V any](k K, maps ...map[K]V) (V, bool) {
	for _, m := range maps {
		if v, ok := m[k]; ok {
			return v, true
		}
	}
	var zero V
	return zero, false
}
```

## Walkthrough

`Lookup(a, {a:0}, {a:2})` returns `0, true`: the first layer holds the key, zero value or not.

## Pitfalls

- Testing `m[k] != zero` instead of using `ok`, which skips stored zeros.
- Searching in reverse, so the lowest-priority layer wins.
- Returning `false` when a later map has the key.
