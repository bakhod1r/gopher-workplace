// Package minval — Gopher Workplace challenge.
package minval

import "math"

// Stats holds a collection of numbers.
type Stats struct {
	Values []float64
}

// Min returns the smallest value. Returns +Inf for empty stats.
//
// Examples:
//
//	Stats{Values: []float64{3, 1, 2}}.Min() => 1
//	Stats{}.Min()                             => +Inf
func (s Stats) Min() float64 {
	// TODO(candidate): implement this.
	_ = math.Inf // hint
	panic("not implemented")
}
