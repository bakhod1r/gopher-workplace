// Package gcpausesum — Gopher Workplace challenge.
package gcpausesum

// Total returns the sum of the pauses, ignoring negative entries.
//
// Examples:
//
//	Total([]int64{100, 200}) => 300
func Total(pauses []int64) int64 {
	panic("not implemented")
}

// Worst returns the longest pause and its index, earliest index on a tie, and
// -1 when there are no valid pauses.
//
// Examples:
//
//	Worst([]int64{100, 500, 200}) => 500, 1
func Worst(pauses []int64) (int64, int) {
	panic("not implemented")
}

// FractionOf returns what share of a wall-clock window the pauses consumed —
// the number that matters, because 5ms of pause in a second is invisible and
// 5ms in every 10ms request is the whole latency budget. A non-positive
// window gives 0.
//
// Examples:
//
//	FractionOf([]int64{5_000_000}, 1_000_000_000) => 0.005
func FractionOf(pauses []int64, windowNS int64) float64 {
	panic("not implemented")
}

// WithinBudget reports whether both conditions hold: the pauses take at most
// maxFraction of the window, and no single pause exceeds maxPauseNS. Both
// limits are inclusive.
//
// Examples:
//
//	WithinBudget([]int64{1_000_000}, 1_000_000_000, 0.01, 5_000_000) => true
func WithinBudget(pauses []int64, windowNS int64, maxFraction float64, maxPauseNS int64) bool {
	panic("not implemented")
}
