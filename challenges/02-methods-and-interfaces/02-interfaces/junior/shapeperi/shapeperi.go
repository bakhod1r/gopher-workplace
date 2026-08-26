// Package shapeperi — Gopher Workplace challenge.
package shapeperi

import "math"

// Shape is an interface for geometric shapes with perimeters.
type Shape interface {
	Perimeter() float64
}

// Circle has a radius.
type Circle struct{ Radius float64 }

// Perimeter returns the circumference: 2 * π * r.
func (c Circle) Perimeter() float64 {
	_ = math.Pi
	// TODO(candidate): return 2 * math.Pi * c.Radius
	panic("not implemented")
}

// Square has a side.
type Square struct{ Side float64 }

// Perimeter returns 4 * side.
func (s Square) Perimeter() float64 {
	// TODO(candidate): return 4 * s.Side
	panic("not implemented")
}
