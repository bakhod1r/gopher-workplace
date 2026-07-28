# Transposing a matrix

## The idea

Element `[i][j]` moves to `[j][i]`. Allocate the transposed shape, then copy:

```go
out := make([][]int, cols)
for j := range out { out[j] = make([]int, rows) }
for i := range grid {
	for j := range grid[i] { out[j][i] = grid[i][j] }
}
```

## Why it matters

Pivoting tables, swapping axes for column-major consumers, and matrix math all
transpose. It cements 2-D indexing and pre-allocation of nested slices.

## Watch out

- Assumes rectangular input; ragged rows need care.
- Allocate each inner row with `make`, or you index into nil.
- Guard the empty grid before reading `grid[0]`.
