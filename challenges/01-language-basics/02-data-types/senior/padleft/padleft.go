// Package padleft left-pads a string to a width with a fill byte.
// A planted bug pads on the wrong side.
package padleft

import "strings"

// Pad returns s left-padded with fill to at least width runes. If s is already
// >= width, it is returned unchanged.
func Pad(s string, width int, fill byte) string {
	n := width - len([]rune(s))
	if n <= 0 {
		return s
	}
	pad := strings.Repeat(string(fill), n)
	// CHANGE CODE BELOW THIS LINE
	return s + pad
	// CHANGE CODE ABOVE THIS LINE
}
