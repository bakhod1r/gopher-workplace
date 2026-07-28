// Package clockmod does 24-hour clock arithmetic.
package clockmod

// AddHours returns the hour of day after adding add hours to h (0..23).
// add may be negative or large; the result is always in 0..23.
//
// TODO(candidate): implement with modulo that never returns negative.
func AddHours(h, add int) int {
	panic("not implemented")
}
