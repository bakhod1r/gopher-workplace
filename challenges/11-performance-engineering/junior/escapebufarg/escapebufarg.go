// Package escapebufarg — Gopher Workplace challenge.
package escapebufarg

// Point is a small value type.
type Point struct{ X, Y int }

// Add returns the component-wise sum. Returning a small struct by value keeps
// it on the caller's stack, so this must not allocate.
//
// Examples:
//
//	Add(Point{1, 2}, Point{3, 4}) => Point{4, 6}
func Add(a, b Point) Point {
	panic("not implemented")
}

// AddInto writes the component-wise sum through dst instead of returning a
// newly allocated point. It must not allocate either.
//
// Examples:
//
//	AddInto(&p, Point{1, 2}, Point{3, 4}) sets p to Point{4, 6}
func AddInto(dst *Point, a, b Point) {
	panic("not implemented")
}
