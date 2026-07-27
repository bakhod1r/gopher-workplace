// Package point — Gopher Workplace challenge.
package point

// Point is a 2-D coordinate. It is a struct value: copying it copies its fields.
type Point struct {
	X int
	Y int
}

// Translate returns a new Point shifted by (dx, dy). It must not modify the
// caller's p — build and return a fresh Point.
//
// Examples:
//
//	Translate(Point{1, 2}, 3, 4)   => Point{4, 6}
//	Translate(Point{0, 0}, -1, -1) => Point{-1, -1}
//	Translate(Point{5, 5}, 0, 0)   => Point{5, 5}
func Translate(p Point, dx, dy int) Point {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
