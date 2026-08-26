// Package fmtstringer — Gopher Workplace challenge.
package fmtstringer

import "fmt"

// Point is a 2D point.
type Point struct {
	X, Y int
}

// String implements fmt.Stringer to format Point as "(X,Y)".
//
// Examples:
//
//	Point{1, 2}.String() => "(1,2)"
func (p Point) String() string {
	// TODO(candidate): format as "(X,Y)".
	_ = fmt.Sprintf // hint
	panic("not implemented")
}
