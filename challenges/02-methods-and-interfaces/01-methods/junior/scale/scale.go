// Package scale — Gopher Workplace challenge.
package scale

// Vector represents a 2D vector.
type Vector struct {
	X, Y float64
}

// Scale multiplies both components by factor, mutating the vector in place.
//
// Examples:
//
//	v := Vector{3, 4}; v.Scale(2) // v is now {6, 8}
//	v := Vector{1, -1}; v.Scale(0) // v is now {0, 0}
func (v *Vector) Scale(factor float64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
