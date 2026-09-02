# Rows That Are All The Same Row

## Intuition

`slices.Repeat` builds a new outer slice and copies the element *values* into it. For `[][]T` an element is a slice header, so all `n` rows carry the same pointer — one array, shared by every row and by `proto`.

## Approach

1. Return an empty result for a non-positive count.
2. Allocate the outer slice of length n.
3. Give each row its own `slices.Clone(proto)`.

## Solution

```go
func Blank[T any](proto []T, n int) [][]T {
	if n <= 0 {
		return [][]T{}
	}
	out := make([][]T, n)
	for i := range out {
		out[i] = slices.Clone(proto)
	}
	return out
}
```

## Walkthrough

`b := Blank(proto, 3); b[0][0] = 7` sets `b[1][0]` and `b[2][0]` to 7 as well, and `proto[0]` with them — three headers, one array.

## Pitfalls

- The same trap in `make([][]T, n)` followed by filling every row with the same slice.
- Reaching for `slices.Repeat` on the inner slice instead — that one is fine, because `T` there is not a reference.
