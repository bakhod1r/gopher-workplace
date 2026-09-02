// Package tempavg — Gopher Workplace challenge.
package tempavg

// AverageReading drains the window's readings and returns their mean.
// A window with no readings returns 0, false.
//
// Examples:
//
//	AverageReading(chan 1,2,3) => 2, true
//	AverageReading(closed empty) => 0, false
//	AverageReading(chan 5) => 5, true
func AverageReading(readings <-chan float64) (float64, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}
