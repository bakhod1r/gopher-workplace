// Package temp uses typed constants for temperature limits.
package temp

// Celsius is a temperature in degrees Celsius.
type Celsius float64

// AbsoluteZero and Boiling as typed Celsius constants.
//
// TODO(candidate): define AbsoluteZero=-273.15, Boiling=100.
const (
	AbsoluteZero Celsius = 0
	Boiling      Celsius = 0
)

// Valid reports whether c is at or above absolute zero.
//
// TODO(candidate): implement.
func Valid(c Celsius) bool {
	panic("not implemented")
}
