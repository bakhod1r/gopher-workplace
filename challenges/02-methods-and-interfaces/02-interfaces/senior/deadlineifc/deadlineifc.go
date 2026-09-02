// Package deadlineifc — Gopher Workplace challenge.
package deadlineifc

import "time"

// Clock reports the current time.
type Clock interface {
	Now() time.Time
}

// FakeClock advances only when told to.
type FakeClock struct {
	T time.Time
}

// Now returns the current fake time.
func (f *FakeClock) Now() time.Time { return f.T }

// Advance moves the clock forward.
func (f *FakeClock) Advance(d time.Duration) { f.T = f.T.Add(d) }

// Op is one unit of work.
type Op interface {
	Do()
}

// CountingOp counts its runs and advances a clock by Cost per run.
type CountingOp struct {
	Clock *FakeClock
	Cost  time.Duration
	Runs  int
}

// Do performs one unit of work.
func (c *CountingOp) Do() {
	// TODO(candidate): count the run and advance the clock by Cost.
	panic("not implemented")
}

// RunUntil runs op repeatedly while the clock is before the deadline and
// returns how many operations completed.
//
// Examples:
//
//	a passed deadline => 0
func RunUntil(clock Clock, deadline time.Time, op Op) int {
	// TODO(candidate): check the deadline before each run.
	panic("not implemented")
}
