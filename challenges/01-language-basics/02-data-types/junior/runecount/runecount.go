// Package runecount — Gopher Workplace challenge.
package runecount

// Count returns the number of characters (runes) in s, not the number of bytes.
// Go strings are UTF-8: a multi-byte character like "é" or "日" is several bytes
// but a single rune, so len(s) (a byte count) is the wrong answer.
//
// Examples:
//
//	Count("abc")   => 3
//	Count("héllo") => 5   // é is 2 bytes, still 1 rune
//	Count("日本")   => 2   // 3 bytes each, 2 runes
//	Count("")      => 0
func Count(s string) int {
	// TODO(candidate): implement this from scratch so all tests pass.
	panic("not implemented")
}
