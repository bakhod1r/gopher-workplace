// Package temperature — Gopher Workplace challenge.
package temperature

// freezingF is the Fahrenheit value of 0 °C.
const freezingF = 32

// CToF converts a Celsius temperature to Fahrenheit: F = C*(9/5) + 32.
//
// Examples:
//
//	CToF(0)   => 32
//	CToF(100) => 212
//	CToF(-40) => -40
func CToF(c float64) float64 {
	// TODO(candidate): implement this. Mind that 9/5 with integer operands
	// truncates to 1 — use a floating-point constant so the ratio is 1.8.
	panic("not implemented")
}
