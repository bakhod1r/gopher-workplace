// Package rectarea computes rectangle geometry via a struct.
package rectarea

// Rect is an axis-aligned rectangle by width and height.
type Rect struct {
	W, H int
}

// Area returns the rectangle's area.
//
// TODO(candidate): implement.
func (r Rect) Area() int {
	panic("not implemented")
}

// Scale returns a new Rect scaled by factor (fields multiplied).
//
// TODO(candidate): implement, returning a new Rect.
func (r Rect) Scale(factor int) Rect {
	panic("not implemented")
}
