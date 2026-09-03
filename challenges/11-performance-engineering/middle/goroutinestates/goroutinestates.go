// Package goroutinestates — Gopher Workplace challenge.
package goroutinestates

// G is one goroutine from a stack dump: its state and the function it is
// waiting in.
type G struct {
	State string
	Top   string
}

// Count tallies the goroutines by state, the first thing to look at in a
// dump: thousands in "chan receive" is a leak, thousands "runnable" is a
// scheduler that cannot keep up. Goroutines with an empty state are counted
// as "unknown".
//
// Examples:
//
//	Count([{running x} {chan receive y}]) => {"running":1, "chan receive":1}
func Count(gs []G) map[string]int {
	panic("not implemented")
}

// Blocked reports how many goroutines are in a state that is not making
// progress: anything other than "running" and "runnable".
//
// Examples:
//
//	Blocked([{running x} {chan receive y}]) => 1
func Blocked(gs []G) int {
	panic("not implemented")
}

// LeakSuspects returns the wait sites where at least threshold goroutines are
// blocked in the same state, ordered by count descending, then state
// ascending, then site ascending — the shape of a goroutine leak in a dump. A
// threshold below 1 is treated as 1.
//
// Examples:
//
//	LeakSuspects(gs, 100) => sites with 100+ goroutines stacked up
func LeakSuspects(gs []G, threshold int) []string {
	panic("not implemented")
}
