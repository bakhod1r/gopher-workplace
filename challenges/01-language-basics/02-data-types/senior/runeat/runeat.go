// Package runeat returns the n-th rune of a string.
// A planted bug indexes bytes instead of runes.
package runeat

// At returns the n-th rune (0-based) of s, counting characters not bytes.
// Returns (0,false) if n is out of range.
func At(s string, n int) (rune, bool) {
	rs := []rune(s)
	// CHANGE CODE BELOW THIS LINE
	if n < 0 || n >= len(s) {
		// CHANGE CODE ABOVE THIS LINE
		return 0, false
	}
	return rs[n], true
}
