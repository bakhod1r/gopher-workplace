# 2-D grids

## Intuition

A 2-D structure is a slice of slices, `[][]T`, indexed `grid[row][col]` in
row-major order. Rows can be independent (and even ragged):

```go
for r := range grid { for c := range grid[r] { use(grid[r][c]) } }
```

## Approach

1. Declare a zero-valued [7][10]bool (all false by default).
2. Range over taken, reading r=seat[0], c=seat[1].
3. Bounds-check 0<=r<7 and 0<=c<10; only then set grid[r][c]=true.
4. Return the grid.

## Solution

```go
func SeatMap(taken [][2]int) [7][10]bool {
	var grid [7][10]bool
	for _, seat := range taken {
		r, c := seat[0], seat[1]
		if r >= 0 && r < 7 && c >= 0 && c < 10 {
			grid[r][c] = true
		}
	}
	return grid
}
```

## Walkthrough

For {{1,2},{6,9}}: {1,2} is in range -> grid[1][2]=true; {6,9} in range -> grid[6][9]=true. All other cells stay at their false zero value.

## Pitfalls

- Allocate each row with its own `make`; assigning one row to all aliases them.
- `grid[row][col]` — outer is row, inner is column.
- Ragged rows are allowed; don't assume a fixed width.
