// Package ifaceallocount — Gopher Workplace challenge.
package ifaceallocount

// Shape is satisfied by anything that can report its area.
type Shape interface{ Area() int }

// Rect is a rectangle. Its Area method has a pointer receiver, so only a
// *Rect satisfies Shape.
type Rect struct{ W, H int }

// Area returns the rectangle's area.
//
// Examples:
//
//	(&Rect{2, 3}).Area() => 6
func (r *Rect) Area() int {
	panic("not implemented")
}

// TotalArea sums the areas of the shapes. Ranging over an existing slice of
// interfaces must not allocate.
//
// Examples:
//
//	TotalArea([]Shape{&Rect{2, 3}}) => 6
func TotalArea(shapes []Shape) int {
	panic("not implemented")
}
