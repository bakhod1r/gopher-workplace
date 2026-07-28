// Package runecount counts characters (runes) in a UTF-8 string. A planted bug
// uses len(s), which counts BYTES, over-counting multibyte characters.
package runecount

// CharCount returns the number of Unicode code points in s.
func CharCount(s string) int {
	// CHANGE CODE BELOW THIS LINE
	return len(s)
	// CHANGE CODE ABOVE THIS LINE
}
