# Clone A Map

## Intuition

A map copy is shallow just like a slice clone, so a map of slices still hands out shared backing arrays. Independence stops at the top level.

## Approach

1. Clone the map.
2. Replace a nil clone with an allocated empty map.
3. Return it.

## Solution

```go
func Snapshot[K comparable, V any](m map[K]V) map[K]V {
	out := maps.Clone(m)
	if out == nil {
		out = make(map[K]V)
	}
	return out
}
```

## Walkthrough

Deleting a key from the snapshot leaves the live map untouched, because the two maps have separate storage.

## Pitfalls

- Returning `m` itself, so writers corrupt the snapshot.
- Returning a nil map, which panics on the caller's first write.
- Assuming a cloned `map[string][]int` deep-copies the slices.
