// Package countwords counts whitespace-separated words.
// A planted bug counts empty tokens produced by runs of spaces.
package countwords

import "strings"

// Count returns the number of words in s (maximal runs of non-space).
// Multiple/leading/trailing spaces must not create empty words.
func Count(s string) int {
	// CHANGE CODE BELOW THIS LINE
	return len(strings.Split(s, " "))
	// CHANGE CODE ABOVE THIS LINE
}
