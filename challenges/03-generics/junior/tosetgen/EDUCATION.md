# To Set

## Intuition

A map already enforces key uniqueness, so the loop can assign unconditionally. `struct{}{}` stores nothing, making the map a pure set.

## Approach

1. Allocate the map with capacity `len(s)`.
2. Assign `out[e] = struct{}{}` for each element.
3. Return `out`.

## Solution

```go
func ToSet[T comparable](s []T) map[T]struct{} {
	out := make(map[T]struct{}, len(s))
	for _, e := range s {
		out[e] = struct{}{}
	}
	return out
}
```

## Walkthrough

`ToSet([]int{1, 1, 2})` writes key 1 twice — the second write is a no-op — and key 2 once, leaving two keys.

## Pitfalls

- Returning a nil map, which panics on write if the caller adds to it.
- Using `map[T]bool` when the signature asks for `map[T]struct{}`.
- Adding a `seen` check the map already performs.
