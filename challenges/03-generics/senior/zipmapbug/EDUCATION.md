# Zip That Panics On Short Input

## Intuition

Ranging `keys` and indexing `vals` is safe only when the slices are the same length, which is precisely what malformed input violates.

## Approach

1. Take the smaller of the two lengths.
2. Assign pairs by index below that bound.

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

`ZipMap([]string{"a","b"}, []int{1})` indexes `vals[1]` and panics; the fix stops after one pair.

## Pitfalls

- Ranging one slice while indexing another.
- Guarding with a length equality check and returning nil — the spec says truncate.
- Assuming the result size equals `len(keys)`.
