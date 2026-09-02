# What comparable Really Means

## Intuition

`comparable` means `==` is defined and cannot panic. Struct comparability is field-wise, which is why one slice field disqualifies the whole struct.

## Approach

1. Build a `seen` set from the elements.
2. Return `len(seen)`.

## Solution

```go
func CountDistinct[T comparable](s []T) int {
	seen := make(map[T]struct{}, len(s))
	for _, v := range s {
		seen[v] = struct{}{}
	}
	return len(seen)
}
```

## Walkthrough

`CountDistinct([]Point{{1,2},{1,2}})` hashes two equal structs to the same key, leaving one entry.

## Pitfalls

- Expecting a struct with a `[]string` field to satisfy `comparable`.
- Confusing `comparable` with `cmp.Ordered` — the latter also needs `<`.
- Counting `len(s)` instead of the distinct values.
