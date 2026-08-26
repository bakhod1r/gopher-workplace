// Package maxval — Gopher Workplace challenge.
package maxval

import "math"

// Stats holds a collection of numbers.
type Stats struct {
	Values []float64
}

// Max returns the largest value. Returns -Inf for empty stats.
//
// Examples:
//
//	Stats{Values: []float64{3, 1, 2}}.Max() => 3
//	Stats{}.Max()                             => -Inf
func (s Stats) Max() float64 {
	// TODO(candidate): implement this.
	_ = math.Inf // hint
	panic("not implemented")
}
