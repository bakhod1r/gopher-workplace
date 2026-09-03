// Package byvalue — Gopher Workplace challenge.
package byvalue

// Point is a two-dimensional integer point.
type Point struct {
	X, Y int
}

// Scale returns p with both coordinates multiplied by f.
//
// Nothing about the result outlives the call, so nothing needs the heap.
//
// Examples:
//
//	Scale(Point{2, 3}, 2) => Point{4, 6}
func Scale(p Point, f int) Point {
	panic("not implemented")
}
