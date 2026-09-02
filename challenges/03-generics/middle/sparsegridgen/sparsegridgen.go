// Package sparsegridgen — Gopher Workplace challenge.
package sparsegridgen

// point is a grid coordinate.
type point struct {
	X int
	Y int
}

// Grid is a sparse two-dimensional grid of T.
// Use NewGrid to create one.
type Grid[T any] struct {
	cells map[point]T
	def   T
}

// NewGrid returns a sparse grid whose unset cells read as def.
func NewGrid[T any](def T) *Grid[T] {
	// TODO(candidate): store the default and allocate the cell map.
	panic("not implemented")
}

// Set stores v at (x, y).
func (g *Grid[T]) Set(x, y int, v T) {
	// TODO(candidate): store the value at the coordinate.
	panic("not implemented")
}

// At returns the value at (x, y), or the default.
func (g *Grid[T]) At(x, y int) T {
	// TODO(candidate): look the coordinate up, falling back to the default.
	panic("not implemented")
}

// Filled returns how many cells have been set.
func (g *Grid[T]) Filled() int {
	// TODO(candidate): report the number of stored cells.
	panic("not implemented")
}
