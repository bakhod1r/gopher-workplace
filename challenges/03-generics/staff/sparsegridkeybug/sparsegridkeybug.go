// Package sparsegridkeybug — Gopher Workplace challenge.
package sparsegridkeybug

// Point is a two-dimensional integer coordinate.
type Point struct {
	X, Y int
}

// Grid is a sparse two-dimensional grid of values.
type Grid[T any] struct {
	cells map[int64]T
}

// gridKey packs a point into a single map key.
// Distinct points must produce distinct keys for coordinates that fit in an int32.
func gridKey(p Point) int64 {
	// CHANGE CODE BELOW THIS LINE
	return int64(p.X + p.Y)
	// CHANGE CODE ABOVE THIS LINE
}

// Set stores v at p.
func (g *Grid[T]) Set(p Point, v T) {
	if g.cells == nil {
		g.cells = make(map[int64]T)
	}
	g.cells[gridKey(p)] = v
}

// Get returns the value stored at p.
func (g *Grid[T]) Get(p Point) (T, bool) {
	v, ok := g.cells[gridKey(p)]
	return v, ok
}

// Len reports how many cells hold a value.
func (g *Grid[T]) Len() int {
	return len(g.cells)
}
