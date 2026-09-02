# Flatten

## Intuition

The first pass costs nothing but avoids repeated reallocation in the second, which matters once the groups are large.

## Approach

1. Sum `len(g)` over all groups into `n`.
2. Allocate `out` with capacity `n`.
3. Append every group with `append(out, g...)`.

## Solution

```go
func Flatten[T any](groups [][]T) []T {
	n := 0
	for _, g := range groups {
		n += len(g)
	}
	out := make([]T, 0, n)
	for _, g := range groups {
		out = append(out, g...)
	}
	return out
}
```

## Walkthrough

`Flatten([][]int{{1, 2}, {3}})` sizes `out` to capacity 3, then splices `[1 2]` and `[3]`, so no regrowth ever happens.

## Pitfalls

- Appending `g` instead of `g...`, which does not compile.
- Returning the first group when there is only one — the result must be a new slice.
- Special-casing empty groups; `append(out)` with no elements is already a no-op.
