// Package commaok looks up a score, distinguishing "absent" from "zero". A
// planted bug ignores the comma-ok boolean, so a missing key (zero value) is
// reported as a real 0 score.
package commaok

// Lookup returns the score for name and whether it was present.
func Lookup(scores map[string]int, name string) (score int, ok bool) {
	// CHANGE CODE BELOW THIS LINE
	score = scores[name]
	ok = true
	// CHANGE CODE ABOVE THIS LINE
	return
}
