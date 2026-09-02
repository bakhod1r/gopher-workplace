// Package apithrottle — Gopher Workplace challenge.
package apithrottle

// PeakInFlight issues every request concurrently behind a semaphore of
// capacity limit and reports the highest number of requests that were ever in
// flight at once, which never exceeds limit. limit is >= 1.
//
// Examples:
//
//	PeakInFlight(5 requests, 2, do)  => a value in 1..2
//	PeakInFlight(3 requests, 1, do)  => 1
//	PeakInFlight(nil, 4, do)         => 0
func PeakInFlight(requests []string, limit int, do func(string)) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
