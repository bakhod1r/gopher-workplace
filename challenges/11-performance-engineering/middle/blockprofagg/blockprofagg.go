// Package blockprofagg — Gopher Workplace challenge.
package blockprofagg

// Event is one sampled blocking event from a block profile: where it
// happened and how long the goroutine waited, in nanoseconds.
type Event struct {
	Site string
	Wait int64
}

// Scale converts a sampled block profile back to whole-program numbers.
// runtime.SetBlockProfileRate(rate) records roughly one event per rate
// nanoseconds of blocking, so each sampled wait stands for rate times as much
// real blocking. A rate at or below 1 records everything and scales by 1.
//
// Examples:
//
//	Scale(100, 1_000_000) => 100_000_000
func Scale(wait int64, rate int) int64 {
	panic("not implemented")
}

// Totals aggregates the events per site, scaling each wait by the profile
// rate, and returns the per-site totals. Events with a non-positive wait are
// dropped.
//
// Examples:
//
//	Totals([{a 100}, {a 50}], 1) => {"a": 150}
func Totals(events []Event, rate int) map[string]int64 {
	panic("not implemented")
}

// FractionBlocked returns what share of a wall-clock window, across all
// goroutines, was spent blocked: the scaled total over windowNS. A
// non-positive window gives 0. The result may exceed 1, because many
// goroutines can block at once.
//
// Examples:
//
//	FractionBlocked([{a 50}], 1, 100) => 0.5
func FractionBlocked(events []Event, rate int, windowNS int64) float64 {
	panic("not implemented")
}
