// Package checktemp — Gopher Workplace challenge.
package checktemp

import "errors"

// Sensor range faults.
var (
	ErrBelowRange = errors.New("temperature below sensor range")
	ErrAboveRange = errors.New("temperature above sensor range")
)

// CheckTemp reports whether c is inside the sensor range [-40, 85].
//
// Examples:
//
//	CheckTemp(20)  => nil
//	CheckTemp(100) => ErrAboveRange
func CheckTemp(c float64) error {
	// TODO(candidate): implement this.
	panic("not implemented")
}
