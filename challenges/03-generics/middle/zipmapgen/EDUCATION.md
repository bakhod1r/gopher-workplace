# Zip To Map

## Intuition

Because map assignment overwrites, duplicate headers silently collapse — worth documenting rather than guarding, since callers usually want last-wins.

## Approach

1. Compute the smaller length.
2. Allocate the map with that capacity.
3. Assign each pair by index.

## Solution

```go
func ZipMap[K comparable, V any](keys []K, vals []V) map[K]V {
	n := len(keys)
	if len(vals) < n {
		n = len(vals)
	}
	out := make(map[K]V, n)
	for i := 0; i < n; i++ {
		out[keys[i]] = vals[i]
	}
	return out
}
```

## Walkthrough

`ZipMap([]string{"a","a"}, []int{1,2})` writes `a:1` then overwrites it with `a:2`.

## Pitfalls

- Ranging over `keys` and indexing `vals`, which panics when `vals` is shorter.
- Returning a nil map for empty input.
- Assuming the map's size equals the number of pairs.
