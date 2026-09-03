// Package reportallocsscopebug — Gopher Workplace challenge.
package reportallocsscopebug

// Run models a benchmark run: warmup iterations happen before the timer and
// the allocation counters are reset, and measured iterations are the ones the
// report is about.
type Run struct {
	Warmup   int64
	Measured int64
	Bytes    uint64 // bytes allocated during the measured iterations only
	Allocs   uint64 // allocations during the measured iterations only
}

// PerOp reports the B/op and allocs/op columns for a run: the measured totals
// divided by the number of iterations they cover. A run with no measured
// iterations reports zeros.
//
// Examples:
//
//	PerOp(Run{Warmup: 90, Measured: 10, Bytes: 800, Allocs: 20}) => 80, 2
func PerOp(r Run) (bytesPerOp, allocsPerOp uint64) {
	if r.Measured <= 0 {
		return 0, 0
	}
	// CHANGE CODE BELOW THIS LINE
	n := uint64(r.Warmup + r.Measured)
	// CHANGE CODE ABOVE THIS LINE
	return r.Bytes / n, r.Allocs / n
}
