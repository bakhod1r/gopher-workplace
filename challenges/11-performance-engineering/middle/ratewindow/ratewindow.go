// Package ratewindow — Gopher Workplace challenge.
package ratewindow

// Event is one observation at a monotonically non-decreasing timestamp in
// nanoseconds.
type Event struct {
	NS    int64
	Value int64
}

// CountIn returns how many events fall in the half-open window
// [fromNS, fromNS+widthNS). Events are sorted by NS ascending; use binary
// searches rather than a scan. A non-positive width counts nothing.
//
// Examples:
//
//	CountIn(events, 0, 100) => events with NS in [0,100)
func CountIn(events []Event, fromNS, widthNS int64) int {
	panic("not implemented")
}

// RatePerSec returns the events-per-second rate in that window: the count
// divided by the width in seconds. A non-positive width gives 0.
//
// Examples:
//
//	RatePerSec(events, 0, 1_000_000_000) => events in the first second
func RatePerSec(events []Event, fromNS, widthNS int64) float64 {
	panic("not implemented")
}

// SumIn returns the total Value of the events in the window, which is what
// turns a rate of requests into a rate of bytes.
//
// Examples:
//
//	SumIn(events, 0, 100) => summed Value
func SumIn(events []Event, fromNS, widthNS int64) int64 {
	panic("not implemented")
}
