// Package runetruncate — Gopher Workplace challenge.
package runetruncate

import "unicode/utf8"

// Truncate returns the longest prefix of s that is at most n bytes and
// does not end in the middle of a UTF-8 character.
//
// The result is a substring, so nothing is copied.
//
// Examples:
//
//	Truncate("héllo", 3) => "hé"
func Truncate(s string, n int) string {
	panic("not implemented")
}
