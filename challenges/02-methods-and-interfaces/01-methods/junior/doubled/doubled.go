// Package doubled — Gopher Workplace challenge.
package doubled

// MyFloat is a floating-point number with methods.
type MyFloat float64

// Double returns the value multiplied by 2.
// It does NOT mutate the receiver — uses a value receiver.
//
// Examples:
//
//	MyFloat(3.5).Double() => 7.0
//	MyFloat(0).Double()   => 0
func (f MyFloat) Double() MyFloat {
	// TODO(candidate): implement this.
	panic("not implemented")
}
