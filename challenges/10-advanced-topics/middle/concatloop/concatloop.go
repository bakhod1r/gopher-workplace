// Package concatloop — Gopher Workplace challenge.
package concatloop

import "strings"

// Join concatenates parts end to end.
//
// Strings are immutable, so += allocates a new string and copies both sides
// every round — quadratic in the total length.
//
// Examples:
//
//	Join([]string{"a", "bc"}) => "abc"
func Join(parts []string) string {
	// CHANGE CODE BELOW THIS LINE
	out := ""
	for _, p := range parts {
		out += p
	}
	return out
	// CHANGE CODE ABOVE THIS LINE
}
