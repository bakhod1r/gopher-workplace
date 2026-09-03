// Package pointerwalk — Gopher Workplace challenge.
package pointerwalk

import "unsafe"

// SumInt32 totals n consecutive int32 values starting at p.
//
// This is the shape a C API hands you: a pointer and a count, with no
// slice. n <= 0 or a nil pointer totals 0.
//
// Examples:
//
//	SumInt32(&a[0], 3) => a[0] + a[1] + a[2]
func SumInt32(p *int32, n int) int64 {
	panic("not implemented")
}
