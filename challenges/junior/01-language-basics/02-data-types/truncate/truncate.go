// Package truncate — Gopher Workplace challenge.
package truncate

// WholePart returns the whole-dollar part of a monetary amount by converting a
// float64 to an int. A float64→int conversion in Go truncates toward zero (it
// drops the fraction, it does not round), so 9.99 yields 9 and -9.99 yields -9.
//
// Examples:
//
//	WholePart(9.99)  => 9
//	WholePart(-9.99) => -9
//	WholePart(4.0)   => 4
func WholePart(amount float64) int {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
