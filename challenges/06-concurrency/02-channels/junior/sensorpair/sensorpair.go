// Package sensorpair — Gopher Workplace challenge.
package sensorpair

// CombinedReading receives exactly one value from each sensor channel using
// select, in whichever order they become ready, and returns their sum.
//
// Examples:
//
//	CombinedReading(chan 21, chan 40) => 61
//	CombinedReading(chan 0, chan 55)  => 55
//	CombinedReading(chan -5, chan 5)  => 0
func CombinedReading(temp, humidity <-chan int) int {
	// TODO(candidate): implement this.
	panic("not implemented")
}
