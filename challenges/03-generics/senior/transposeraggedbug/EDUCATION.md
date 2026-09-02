# Transpose That Assumes A Rectangle

## Intuition

Taking the width from `m[0]` drops every column beyond it, so a matrix whose first row is short loses data.

## Approach

1. Scan for the longest row.
2. Build that many columns.
3. Skip rows that are too short for the current column.

## Solution

```go
func Transpose[T any](m [][]T) [][]T {
	width := 0
	for _, row := range m {
		if len(row) > width {
			width = len(row)
		}
	}
	out := make([][]T, width)
	for c := 0; c < width; c++ {
		col := make([]T, 0, len(m))
		for _, row := range m {
			if c < len(row) {
				col = append(col, row[c])
			}
		}
		out[c] = col
	}
	return out
}
```

## Walkthrough

`Transpose([[1],[2,3]])` builds one column and discards the `3`.

## Pitfalls

- Panicking on `m[0]` for an empty input.
- Assuming a ragged transpose round-trips — it only does when no interior row is short.
