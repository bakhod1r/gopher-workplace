// Package hexdigit maps a nibble to its lowercase hex character.
package hexdigit

// Digit returns the lowercase hex character for n (0..15): '0'..'9' then
// 'a'..'f'. For n outside 0..15 it returns '?'.
//
// TODO(candidate): implement using byte/rune arithmetic.
func Digit(n int) byte {
	panic("not implemented")
}
