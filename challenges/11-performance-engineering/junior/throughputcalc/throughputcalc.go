// Package throughputcalc — Gopher Workplace challenge.
package throughputcalc

// OpsPerSec converts a ns/op measurement into operations per second on one
// core. A non-positive input gives 0.
//
// Examples:
//
//	OpsPerSec(1000) => 1e6
func OpsPerSec(nsPerOp float64) float64 {
	panic("not implemented")
}

// Capacity scales a single-core rate across cores, assuming perfect scaling,
// and returns the whole operations per second the machine could serve. A
// non-positive input gives 0.
//
// Examples:
//
//	Capacity(1000, 8) => 8000000
func Capacity(nsPerOp float64, cores int) int64 {
	panic("not implemented")
}
