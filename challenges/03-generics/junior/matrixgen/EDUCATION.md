# Generic Matrix

## Intuition

A single allocation keeps the cells contiguous and makes the bounds check explicit — with a slice of slices you get two allocations per row and two panics to guard.

## Approach

1. `NewMatrix`: clamp negatives, allocate `rows*cols` cells.
2. `At`: guard all four bounds, then read the flat index.
3. `Set`: same guard, then write, returning whether it happened.

## Solution

```go
func NewMatrix[T any](rows, cols int) *Matrix[T] {
	if rows < 0 {
		rows = 0
	}
	if cols < 0 {
		cols = 0
	}
	return &Matrix[T]{rows: rows, cols: cols, cells: make([]T, rows*cols)}
}

func (m *Matrix[T]) At(r, c int) (T, bool) {
	if r < 0 || c < 0 || r >= m.rows || c >= m.cols {
		var zero T
		return zero, false
	}
	return m.cells[r*m.cols+c], true
}

func (m *Matrix[T]) Set(r, c int, v T) bool {
	if r < 0 || c < 0 || r >= m.rows || c >= m.cols {
		return false
	}
	m.cells[r*m.cols+c] = v
	return true
}
```

## Walkthrough

`Set(1, 1, 5)` on a 2x2 matrix writes flat index `1*2+1 = 3`, and `At(1, 1)` reads the same slot.

## Pitfalls

- Computing `c*m.rows + r`, which transposes the matrix silently.
- Checking only the row bound.
- Allocating `rows` slices of `cols` and forgetting to allocate the inner ones.
