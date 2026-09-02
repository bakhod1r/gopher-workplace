# Cloning Is Shallow

## Intuition

"Deep" here means exactly one level deeper than the stdlib goes, which is the level that bites people using maps of slices.

## Approach

1. Allocate the result map.
2. For each entry, copy the value slice into a fresh slice.
3. Store the copy.

## Solution

```go
func DeepClone[K comparable, V any](m map[K][]V) map[K][]V {
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

Writing into `clone["a"][0]` leaves `m["a"][0]` untouched, unlike with a shallow clone.

## Pitfalls

- Using `maps.Clone` and assuming the values are independent.
- Copying with `out[k] = v[:]`, which still aliases.
- Returning nil for a nil input when the contract promises non-nil.
