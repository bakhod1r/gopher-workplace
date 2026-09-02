// Package sensorlast — Gopher Workplace challenge.
package sensorlast

// LastReading drains the reading stream and returns the final value the
// device sent, plus true. A device that never reported returns 0, false.
//
// Examples:
//
//	LastReading(chan 1,2,3) => 3, true
//	LastReading(closed empty) => 0, false
//	LastReading(chan 7) => 7, true
func LastReading(readings <-chan int) (int, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
