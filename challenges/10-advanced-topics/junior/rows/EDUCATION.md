# One Allocation For The Whole Grid

## Intuition

A `[][]int` is a slice of headers. Nothing says those headers must point at different allocations — carving one flat array into `r` windows gives the same API with one allocation and perfect locality.

## Approach

1. Reject non-positive dimensions.
2. Allocate `flat := make([]int, r*c)` and `out := make([][]int, r)`.
3. Point row `i` at `flat[i*c : (i+1)*c : (i+1)*c]`.

## Solution

```go
// Rows returns an r-by-c grid of zeros whose rows are consecutive
// windows into a single backing array.
//
// Allocating each row separately costs r allocations and scatters the grid
// across the heap; this must cost two.
//
// Examples:
//
// 	Rows(2, 3) => a 2x3 grid, rows 0 and 1 adjacent in memory
func Rows(r, c int) [][]int {
	if r <= 0 || c <= 0 {
		return nil
	}
	flat := make([]int, r*c)
	out := make([][]int, r)
	for i := range out {
		out[i] = flat[i*c : (i+1)*c : (i+1)*c]
	}
	return out
}
```

## Walkthrough

Rows(2,3) allocates six ints. Row 0 is `flat[0:3:3]`, row 1 is `flat[3:6:6]`, so `&g[0][2]` and `&g[1][0]` are neighbours.

## Pitfalls

- Omitting the capacity bound — an `append` to row 0 would overwrite row 1.
- `make([][]int, r)` alone leaves every row nil.
