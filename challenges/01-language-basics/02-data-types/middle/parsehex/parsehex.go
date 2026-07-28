// Package parsehex parses a hex string into an int.
package parsehex

// Parse converts a lowercase/uppercase hex string (no 0x prefix) to an int.
// Returns (0, false) on any non-hex character or empty input.
//
// TODO(candidate): implement digit-by-digit (value 0-15 per char).
func Parse(s string) (int, bool) {
	panic("not implemented")
}
