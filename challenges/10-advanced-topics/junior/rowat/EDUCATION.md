# Reach Into A Slice Of Slices

## Intuition

Indexing is only safe inside the range, and the range is `0` up to `len(g)`. Checking it yourself turns a crash into a value the caller can branch on.

## Approach

1. Return `nil, false` for `i < 0` or `i >= len(g)`.
2. Return `g[i], true`.

## Solution

```go
// Row returns the i-th row of g and whether it exists.
//
// An out-of-range index is a missing row, not a panic. The row is returned
// as a view, so writes through it reach g.
//
// Examples:
//
// 	Row([][]int{{1}, {2}}, 1) => []int{2}, true
func Row(g [][]int, i int) ([]int, bool) {
	if i < 0 || i >= len(g) {
		return nil, false
	}
	return g[i], true
}
```

## Walkthrough

`Row(nil, 0)` sees `len(g) == 0`, so 0 is already out of range and the nil check is not needed separately.

## Pitfalls

- Checking only `i >= len(g)` and letting a negative index panic.
- Copying the row to "be safe", which breaks the caller's writes and allocates.
