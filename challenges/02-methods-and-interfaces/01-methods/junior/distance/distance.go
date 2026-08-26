// Package distance — Gopher Workplace challenge.
package distance

import "math"

// Point represents a 2D point.
type Point struct {
	X, Y float64
}

// DistanceTo returns the Euclidean distance to another point.
//
// Examples:
//
//	Point{0,0}.DistanceTo(Point{3,4}) => 5
//	Point{1,1}.DistanceTo(Point{1,1}) => 0
func (p Point) DistanceTo(other Point) float64 {
	// TODO(candidate): implement this.
	_ = math.Sqrt // hint: you may want this
	panic("not implemented")
}
