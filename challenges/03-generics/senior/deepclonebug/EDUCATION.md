# A Snapshot That Keeps Changing

## Intuition

The cloned map has its own buckets but every value is the same slice header, pointing at the same backing array the collector keeps writing to.

## Approach

1. Allocate the result map.
2. Copy each value slice into a fresh slice.
3. Store the copies.

## Solution

```go
func Snapshot[K comparable, V any](m map[K][]V) map[K][]V {
	out := make(map[K][]V, len(m))
	for k, v := range m {
		cp := make([]V, len(v))
		copy(cp, v)
		out[k] = cp
	}
	return out
}
```

## Walkthrough

Writing `snap["a"][0] = 9` changes `m["a"][0]` too, because both headers address one array.

## Pitfalls

- Trusting `maps.Clone` for maps of slices, maps, or pointers.
- Copying with `out[k] = v[:]`, which still aliases.
- Returning nil for a nil input when the contract promises non-nil.
