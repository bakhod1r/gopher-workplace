# Nested iteration over 2D data

## Intuition

A slice of slices is ragged; nested `range` visits each element without assuming a fixed width.

## Approach

1. Range the rows, then range each row's cells.
2. Accumulate every cell.

## Solution

```go
func GridSum(g [][]int) int {
	total := 0
	for _, row := range g {
		for _, c := range row {
			total += c
		}
	}
	return total
}
```

## Walkthrough

`[[1 2],[3 4]]` sums 1+2+3+4 = 10.

## Pitfalls

- Don't assume all rows share a length; range each row independently.
- A nil or empty grid ranges zero times.
