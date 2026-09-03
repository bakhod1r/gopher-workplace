// Package resettimerskipbug — Gopher Workplace challenge.
package resettimerskipbug

// Timer models the benchmark timer. Add accumulates measured nanoseconds into
// a running total; Reset marks the current position so that everything before
// it is excluded; Elapsed reports only what was measured after the last reset.
type Timer struct {
	total int64
	mark  int64
}

// Add records ns of measured time. Negative durations are ignored.
//
// Examples:
//
//	t.Add(10)
func (t *Timer) Add(ns int64) {
	if ns < 0 {
		return
	}
	t.total += ns
}

// Reset discards everything measured so far, exactly as b.ResetTimer does.
//
// Examples:
//
//	t.Add(500); t.Reset(); t.Elapsed() => 0
func (t *Timer) Reset() {
	// CHANGE CODE BELOW THIS LINE
	t.mark = 0
	// CHANGE CODE ABOVE THIS LINE
}

// Elapsed returns the time measured since the last reset.
//
// Examples:
//
//	t.Elapsed() => 0
func (t *Timer) Elapsed() int64 {
	return t.total - t.mark
}

// Benchmark runs the setup, resets the timer, then records n iterations of
// work, and reports the measured total.
//
// Examples:
//
//	Benchmark(500, 7, 3) => 21
func Benchmark(setupNS, workNS, n int64) int64 {
	var t Timer
	t.Add(setupNS)
	t.Reset()
	for i := int64(0); i < n; i++ {
		t.Add(workNS)
	}
	return t.Elapsed()
}
