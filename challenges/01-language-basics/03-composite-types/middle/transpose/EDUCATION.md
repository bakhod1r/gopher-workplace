# Transposing a matrix

## Intuition

Element `[i][j]` moves to `[j][i]`. Allocate the transposed shape, then copy:

```go
out := make([][]int, cols)
for j := range out { out[j] = make([]int, rows) }
for i := range grid {
	for j := range grid[i] { out[j][i] = grid[i][j] }
}
```

## Approach

1. Empty grid -> empty grid.
2. rows=len(grid), cols=len(grid[0]).
3. Allocate cols x rows.
4. out[j][i]=grid[i][j].
5. Return transposed grid.

## Solution

```go
func Transpose(grid [][]int) [][]int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return [][]int{}
	}
	rows, cols := len(grid), len(grid[0])
	out := make([][]int, cols)
	for j := 0; j < cols; j++ {
		out[j] = make([]int, rows)
		for i := 0; i < rows; i++ {
			out[j][i] = grid[i][j]
		}
	}
	return out
}
```

## Walkthrough

2x3 grid: out[0]=[grid[0][0],grid[1][0]]=[1,4]; out[1]=[2,5]; out[2]=[3,6].

## Pitfalls

- Assumes rectangular input; ragged rows need care.
- Allocate each inner row with `make`, or you index into nil.
- Guard the empty grid before reading `grid[0]`.
