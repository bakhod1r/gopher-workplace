# Filter

## Intuition

Filtering by building a new slice keeps the input untouched and keeps the code obviously correct — in-place deletion while ranging is where bugs live.

## Approach

1. Allocate `out` with capacity `len(s)`.
2. Append each element for which `keep` returns true.
3. Return `out`.

## Solution

```go
func Filter[T any](s []T, keep func(T) bool) []T {
	out := make([]T, 0, len(s))
	for _, e := range s {
		if keep(e) {
			out = append(out, e)
		}
	}
	return out
}
```

## Walkthrough

`Filter([]int{1, 2, 3}, isEven)` rejects 1, keeps 2, rejects 3, and returns `[2]` while the input still reads `[1 2 3]`.

## Pitfalls

- Deleting from `s` while ranging over it, which skips elements.
- Inverting the condition and keeping what should be dropped.
- Returning `nil` when nothing matches, if the tests expect an empty slice.
