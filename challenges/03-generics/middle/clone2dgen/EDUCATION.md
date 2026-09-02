# Cloning A Matrix

## Intuition

A slice of slices is a slice of headers; copying it duplicates the headers while leaving every row pointing at the same array.

## Approach

1. Allocate the outer slice with the row count as capacity.
2. Copy each row into a fresh slice and append it.

## Solution

```go
func Clone2D[T any](m [][]T) [][]T {
	out := make([][]T, 0, len(m))
	for _, row := range m {
		cp := make([]T, len(row))
		copy(cp, row)
		out = append(out, cp)
	}
	return out
}
```

## Walkthrough

Setting `clone[0][0] = 99` leaves `m[0][0]` unchanged, which a shallow clone would not.

## Pitfalls

- Calling `slices.Clone` once and thinking the job is done.
- Appending `row` itself instead of a copy.
- Returning nil for an empty input.
