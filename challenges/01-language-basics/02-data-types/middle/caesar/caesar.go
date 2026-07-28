// Package caesar implements a Caesar shift cipher over ASCII letters.
package caesar

// Shift returns s with each ASCII letter advanced by n places (wrapping within
// its case). Non-letters pass through unchanged. n may be large or negative.
//
// TODO(candidate): implement with rune ranging and modular arithmetic.
func Shift(s string, n int) string {
	panic("not implemented")
}
