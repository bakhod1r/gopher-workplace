# The Clone That Shares Its Rows

## Intuition

The outer `make` gives a new slice of slice headers, and each header is copied — but a header is a pointer, length and capacity. Both matrices end up pointing at the same element arrays.

## Approach

1. Allocate the outer slice.
2. For each row, allocate a slice of the same length and copy the elements in.
3. Store the fresh row.

## Solution

```go
func Clone2D[T any](m [][]T) [][]T {
	out := make([][]T, len(m))
	for i, row := range m {
		r := make([]T, len(row))
		copy(r, row)
		out[i] = r
	}
	return out
}
```

## Walkthrough

`c := Clone2D(m); c[0][0] = 99` also sets `m[0][0]` to 99, because `c[0]` and `m[0]` share one array.

## Pitfalls

- Using `append([]T(nil), row...)` for a nil row and getting `nil` back where an empty slice was expected.
- Stopping at a shallow copy because the tests only ever read from the clone.
