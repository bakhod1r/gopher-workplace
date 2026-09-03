// Package builder — Gopher Workplace challenge.
package builder

import "strings"

// Join concatenates parts separated by sep.
//
// Strings are immutable, so `s += p` allocates a new string every round.
// Build the result in one growing buffer instead.
//
// Examples:
//
//	Join([]string{"a", "b"}, "-") => "a-b"
func Join(parts []string, sep string) string {
	panic("not implemented")
}
