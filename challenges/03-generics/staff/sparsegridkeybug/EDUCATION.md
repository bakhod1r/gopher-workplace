# Coordinates That Collide

## Intuition

Any reduction of two coordinates to one number must be injective. Addition maps a whole diagonal onto a single key, so the grid degenerates from N^2 cells to about 2N. Packing the two halves into the high and low 32 bits keeps them separable.

## Approach

1. Narrow X to 32 bits and shift it into the high half.
2. Narrow Y to 32 bits, reinterpret it as unsigned so sign extension cannot bleed upward, and OR it into the low half.

## Solution

```go
func gridKey(p Point) int64 {
	return int64(int32(p.X))<<32 | int64(uint32(int32(p.Y)))
}

func (g *Grid[T]) Set(p Point, v T) {
	if g.cells == nil {
		g.cells = make(map[int64]T)
	}
	g.cells[gridKey(p)] = v
}

func (g *Grid[T]) Get(p Point) (T, bool) {
	v, ok := g.cells[gridKey(p)]
	return v, ok
}

func (g *Grid[T]) Len() int {
	return len(g.cells)
}
```

## Walkthrough

With the sum as the key, `{1,2}` and `{2,1}` both hash to `3`, so the second `Set` overwrites the first and `Len()` counts one cell instead of two.

## Pitfalls

- Using `x*1000 + y`, which merely moves the collision to larger coordinates.
- Omitting the `uint32` conversion, so a negative Y sign-extends and corrupts the X half.
- Combining with XOR, which makes `{1,2}` and `{2,1}` collide just as badly.
