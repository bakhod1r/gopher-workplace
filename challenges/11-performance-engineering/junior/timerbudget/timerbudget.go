// Package timerbudget — Gopher Workplace challenge.
package timerbudget

// Cost returns the nanoseconds one request spends in an operation that costs
// nsPerOp and is called callsPerRequest times. Non-positive inputs cost 0.
//
// Examples:
//
//	Cost(50, 20) => 1000
func Cost(nsPerOp int64, callsPerRequest int) int64 {
	panic("not implemented")
}

// Headroom returns how many nanoseconds of the budget are left after that
// cost. A negative result means the budget is blown, and the sign is the
// answer, so it must not be clamped.
//
// Examples:
//
//	Headroom(50, 20, 1500) => 500
func Headroom(nsPerOp int64, callsPerRequest int, budgetNS int64) int64 {
	panic("not implemented")
}

// Fits reports whether the cost stays within the budget. Spending exactly the
// budget fits.
//
// Examples:
//
//	Fits(50, 20, 1000) => true
func Fits(nsPerOp int64, callsPerRequest int, budgetNS int64) bool {
	panic("not implemented")
}
