// Package shapearea — Gopher Workplace challenge.
package shapearea

// Shape is an interface for geometric shapes.
type Shape interface {
	Area() float64
}

// Circle has a radius.
type Circle struct {
	Radius float64
}

// Area returns the area of the circle: π * r².
func (c Circle) Area() float64 {
	// TODO(candidate): return math.Pi * c.Radius * c.Radius
	panic("not implemented")
}

// Rectangle has width and height.
type Rectangle struct {
	Width, Height float64
}

// Area returns the area of the rectangle: w * h.
func (r Rectangle) Area() float64 {
	// TODO(candidate): return r.Width * r.Height
	panic("not implemented")
}

// TotalArea sums the areas of all shapes.
func TotalArea(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
