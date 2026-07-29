# Rotation as transpose + reverse

## Intuition

A 90° clockwise rotation of a square matrix decomposes into two simple steps:
transpose (swap `[i][j]` with `[j][i]`), then reverse each row:

```go
// after transposing out:
for i := range out {
	for a, b := 0, n-1; a < b; a, b = a+1, b-1 {
		out[i][a], out[i][b] = out[i][b], out[i][a]
	}
}
```

## Approach

1. Bug: only the transpose (out[i][j]=m[j][i]) runs; the row-reversal step is missing, so the matrix is transposed, not rotated.
2. Fix: after transposing, reverse each row of out (swap ends inward).

## Solution

```go
func Rotate(m [][]int) [][]int {
	n := len(m)
	out := make([][]int, n)
	for i := range out {
		out[i] = make([]int, n)
		for j := 0; j < n; j++ {
			out[i][j] = m[j][i] // transpose
		}
	}
	for i := range out {
		for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
			out[i][l], out[i][r] = out[i][r], out[i][l]
		}
	}
	return out
}
```

## Walkthrough

[[1 2][3 4]] transposes to [[1 3][2 4]]; reversing each row gives [[3 1][4 2]], the 90-degrees-clockwise rotation.

## Pitfalls

- Clockwise = transpose then reverse rows; counter-clockwise = transpose then
  reverse columns.
- The two-pointer reverse stops at the middle.
- An in-place layer-by-layer rotation avoids the second matrix but is trickier.
