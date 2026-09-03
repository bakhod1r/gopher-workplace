// Package nsperop — Gopher Workplace challenge.
package nsperop

// NsPerOp reports the ns/op column: elapsedNS divided by iters, truncated
// toward zero, exactly as the benchmark tool prints it. A non-positive iters
// or a negative elapsedNS reports 0.
//
// Examples:
//
//	NsPerOp(1000, 3) => 333
func NsPerOp(elapsedNS int64, iters int) int64 {
	panic("not implemented")
}

// Faster reports whether candidate is at least pct percent faster than base,
// comparing ns/op values. A non-positive base is never beaten.
//
// Examples:
//
//	Faster(100, 80, 20) => true
func Faster(base, candidate int64, pct float64) bool {
	panic("not implemented")
}
