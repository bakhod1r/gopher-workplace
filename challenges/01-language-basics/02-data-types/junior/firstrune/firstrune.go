// Package firstrune returns the first rune of a string.
package firstrune

// First returns the first rune of s, or the rune 0 (utf8.RuneError is not used)
// for an empty string. Multi-byte runes count as one.
//
// TODO(candidate): implement by ranging over s and returning the first rune.
func First(s string) rune {
	panic("not implemented")
}
