// Package normspaces collapses internal whitespace and trims the ends.
// A planted bug forgets to trim, leaving leading/trailing spaces.
package normspaces

import "strings"

// Normalize collapses every run of whitespace in s to a single space and trims
// leading/trailing whitespace.
func Normalize(s string) string {
	collapsed := strings.Join(strings.Fields(s), " ")
	// CHANGE CODE BELOW THIS LINE
	return " " + collapsed
	// CHANGE CODE ABOVE THIS LINE
}
