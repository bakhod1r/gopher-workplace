// Package runewidth reports the byte width of the first rune in a string.
// A planted bug returns the whole string's byte length.
package runewidth

import "unicode/utf8"

// FirstWidth returns how many bytes the first rune of s occupies (1..4), or 0
// for an empty string.
func FirstWidth(s string) int {
	if s == "" {
		return 0
	}
	_, size := utf8.DecodeRuneInString(s)
	_ = size
	// CHANGE CODE BELOW THIS LINE
	return len(s)
	// CHANGE CODE ABOVE THIS LINE
}
