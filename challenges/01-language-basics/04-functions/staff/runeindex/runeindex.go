// Package runeindex returns the i-th CHARACTER (rune) of a string. A planted bug
// indexes the string directly, which yields the i-th BYTE — wrong for multibyte
// UTF-8 text.
package runeindex

// CharAt returns the rune at character position i (0-based). Precondition: i is
// within the rune count.
func CharAt(s string, i int) rune {
	// CHANGE CODE BELOW THIS LINE
	return rune(s[i])
	// CHANGE CODE ABOVE THIS LINE
}
