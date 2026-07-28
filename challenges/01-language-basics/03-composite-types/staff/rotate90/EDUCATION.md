# Rotation as transpose + reverse

## The idea

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

## Why it matters

Image rotation, matrix transforms, and board games use this. Decomposing a complex
transform into known primitives (transpose, reverse) is a broadly useful technique.

## Watch out

- Clockwise = transpose then reverse rows; counter-clockwise = transpose then
  reverse columns.
- The two-pointer reverse stops at the middle.
- An in-place layer-by-layer rotation avoids the second matrix but is trickier.
