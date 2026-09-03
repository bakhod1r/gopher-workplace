// Package amdahlbound — Gopher Workplace challenge.
package amdahlbound

// MaxSpeedup applies Amdahl's law: optimising the fraction p of the total
// runtime by a factor of s speeds the whole program up by
// 1 / ((1-p) + p/s). A p outside [0,1] or an s below 1 gives 1 — no change.
//
// Examples:
//
//	MaxSpeedup(0.5, 2) => 1.3333333333333333
func MaxSpeedup(p, s float64) float64 {
	panic("not implemented")
}

// Ceiling returns the best speedup any optimisation of that fraction could
// ever reach, which is what happens as s goes to infinity: 1 / (1-p). A p of
// 1 or more has no ceiling and returns math.Inf(1).
//
// Examples:
//
//	Ceiling(0.9) => 10
func Ceiling(p float64) float64 {
	panic("not implemented")
}
