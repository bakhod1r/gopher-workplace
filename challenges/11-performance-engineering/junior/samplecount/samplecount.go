// Package samplecount — Gopher Workplace challenge.
package samplecount

// Sample is one profile sample: how many times this exact stack was observed
// and the nanoseconds each observation represents.
type Sample struct {
	Leaf   string
	Count  int64
	Period int64
}

// Totals reports how a profile's two value columns are derived: the total
// number of samples, and the total nanoseconds, which is the sum of
// Count*Period over every sample. Samples with a non-positive Count or Period
// contribute nothing.
//
// Examples:
//
//	Totals([{"a",3,10},{"b",2,10}]) => 5, 50
func Totals(samples []Sample) (count int64, nanos int64) {
	panic("not implemented")
}
