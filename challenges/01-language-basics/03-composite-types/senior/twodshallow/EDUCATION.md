# Deep-copying nested slices

## The idea

A `[][]int` is a slice of slice-headers. `copy(out, grid)` (or a plain assignment)
duplicates the outer headers but leaves every row pointing at the original backing
arrays. Clone each row:

```go
for i := range grid { out[i] = append([]int{}, grid[i]...) }
```

## Why it matters

Shallow copies of nested structures share their inner data. Mutating the "copy"
corrupts the original — the same shallow-vs-deep lesson as struct slice fields,
one level deeper.

## Watch out

- Each level of nesting needs its own copy.
- `slices.Clone(grid)` is still shallow (rows shared) — clone rows explicitly.
- The cost is O(total elements), unavoidable for true independence.
