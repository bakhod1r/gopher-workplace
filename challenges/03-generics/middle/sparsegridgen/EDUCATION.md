# Sparse Grid

## Intuition

Using a struct key keeps the lookup allocation-free and type-safe, unlike encoding coordinates into a string.

## Approach

1. `Set`: assign at the composite key.
2. `At`: comma-ok lookup falling back to the default.
3. `Filled`: report the map size.

## Solution

```go
func NewGrid[T any](def T) *Grid[T] {
	return &Grid[T]{cells: make(map[point]T), def: def}
}

func (g *Grid[T]) Set(x, y int, v T) {
	g.cells[point{X: x, Y: y}] = v
}

func (g *Grid[T]) At(x, y int) T {
	if v, ok := g.cells[point{X: x, Y: y}]; ok {
		return v
	}
	return g.def
}

func (g *Grid[T]) Filled() int {
	return len(g.cells)
}
```

## Walkthrough

`Set(0,0,def)` stores an entry even though the value equals the default, so `Filled()` reports `1`.

## Pitfalls

- Formatting coordinates into a string key, which allocates on every access.
- Comparing the stored value against the default instead of using `ok`.
- Assuming coordinates are non-negative.
