# Deep-copying nested slices

## Intuition

A `[][]int` is a slice of slice-headers. `copy(out, grid)` (or a plain assignment)
duplicates the outer headers but leaves every row pointing at the original backing
arrays. Clone each row:

```go
for i := range grid { out[i] = append([]int{}, grid[i]...) }
```

## Approach

1. Bug: copy(out, grid) copies only the row-slice headers, so out's rows alias grid's rows. 2. Fix: allocate a fresh row per index: out[i] = append([]int(nil), grid[i]...). 3. Each row is deep-copied, so writes to out don't touch grid.

## Solution

```go
func Clone(grid [][]int) [][]int {
	out := make([][]int, len(grid))
	for i := range grid {
		out[i] = append([]int(nil), grid[i]...)
	}
	return out
}
```

## Walkthrough

copy(out,grid) makes out[0] share grid[0]'s array; out[0][0]=9 mutates grid too. Copying each row into a new array isolates them.

## Pitfalls

- Each level of nesting needs its own copy.
- `slices.Clone(grid)` is still shallow (rows shared) — clone rows explicitly.
- The cost is O(total elements), unavoidable for true independence.
