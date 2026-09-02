// Package latencypeak — Gopher Workplace challenge.
package latencypeak

// PeakLatency drains the sample stream and returns the largest latency seen,
// plus true. A window with no requests returns 0, false.
//
// Examples:
//
//	PeakLatency(chan 30, 90, 40) => 90, true
//	PeakLatency(closed empty)    => 0, false
//	PeakLatency(chan -5, -2)     => -2, true
func PeakLatency(samples <-chan int) (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
