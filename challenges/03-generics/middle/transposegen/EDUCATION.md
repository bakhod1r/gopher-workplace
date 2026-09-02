# Transpose

## Intuition

Ragged input is the only failure mode, so rejecting it once at the top turns the rest of the function into straight-line index arithmetic.

## Approach

1. Return empty for empty input.
2. Reject rows whose length differs from the first.
3. For each column index, collect that position from every row.

## Solution

```go
func Transpose[T any](m [][]T) [][]T {
	out := make([][]T, 0)
	if len(m) == 0 {
		return out
	}
	w := len(m[0])
	for _, row := range m {
		if len(row) != w {
			return make([][]T, 0)
		}
	}
	for c := 0; c < w; c++ {
		col := make([]T, 0, len(m))
		for r := 0; r < len(m); r++ {
			col = append(col, m[r][c])
		}
		out = append(out, col)
	}
	return out
}
```

## Walkthrough

`Transpose([][]int{{1,2},{3,4}})` builds column 0 as `[1 3]` and column 1 as `[2 4]`.

## Pitfalls

- Assuming rectangular input and panicking on a short row.
- Swapping the loop order and producing the original matrix.
- Returning `nil` instead of an empty result on rejection.
