// Package statssum — Gopher Workplace challenge.
package statssum

// Stats holds a collection of numbers.
type Stats struct {
	Values []float64
}

// Sum returns the sum of all values. Empty stats sum to 0.
//
// Examples:
//
//	Stats{Values: []float64{1, 2, 3}}.Sum() => 6
//	Stats{}.Sum()                             => 0
func (s Stats) Sum() float64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
