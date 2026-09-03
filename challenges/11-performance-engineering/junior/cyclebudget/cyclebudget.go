// Package cyclebudget — Gopher Workplace challenge.
package cyclebudget

// CyclesPerOp converts a ns/op measurement into CPU cycles on a machine
// running at ghz gigahertz: at 3 GHz, one nanosecond is three cycles. A
// non-positive input gives 0.
//
// Examples:
//
//	CyclesPerOp(10, 3) => 30
func CyclesPerOp(nsPerOp, ghz float64) float64 {
	panic("not implemented")
}

// Verdict classifies a per-operation cycle count against the rough costs of
// the machine: "register" below 10, "l1" below 100, "memory" below 1000, and
// "syscall" at 1000 or above. A non-positive count is "register".
//
// Examples:
//
//	Verdict(30) => "l1"
func Verdict(cycles float64) string {
	panic("not implemented")
}
