// Package shapearea — Gopher Workplace challenge.
package shapearea

// Shape has a measurable area.
type Shape interface {
	Area() float64
}

// Rect is a rectangle.
type Rect struct {
	W, H float64
}

// Area returns W*H.
//
// Examples:
//
//	Rect{W: 3, H: 4}.Area() => 12
func (r Rect) Area() float64 {
	// TODO(candidate): width times height.
	panic("not implemented")
}

// Square is a square.
type Square struct {
	Side float64
}

// Area returns Side*Side.
func (s Square) Area() float64 {
	// TODO(candidate): side squared.
	panic("not implemented")
}

// TotalArea sums the area of every shape.
func TotalArea(shapes []Shape) float64 {
	// TODO(candidate): accumulate Area().
	panic("not implemented")
}
