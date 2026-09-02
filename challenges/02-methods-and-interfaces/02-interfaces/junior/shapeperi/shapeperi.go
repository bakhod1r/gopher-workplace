// Package shapeperi — Gopher Workplace challenge.
package shapeperi

// Shape has a measurable perimeter.
type Shape interface {
	Perimeter() float64
}

// Rect is a rectangle.
type Rect struct {
	W, H float64
}

// Perimeter returns 2*(W+H).
func (r Rect) Perimeter() float64 {
	// TODO(candidate): 2*(W+H).
	panic("not implemented")
}

// Square is a square.
type Square struct {
	Side float64
}

// Perimeter returns 4*Side.
func (s Square) Perimeter() float64 {
	// TODO(candidate): four sides.
	panic("not implemented")
}

// Circle is a circle of radius R.
type Circle struct {
	R float64
}

// Perimeter returns the circumference.
//
// Examples:
//
//	Circle{R: 1}.Perimeter() => 2*math.Pi
func (c Circle) Perimeter() float64 {
	// TODO(candidate): 2*pi*R.
	panic("not implemented")
}

// LongestPerimeter returns the largest perimeter, or 0 when empty.
func LongestPerimeter(shapes []Shape) float64 {
	// TODO(candidate): track the running maximum.
	panic("not implemented")
}
