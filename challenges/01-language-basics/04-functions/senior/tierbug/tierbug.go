// Package tierbug maps a numeric level to a label with a switch. A planted bug
// adds a stray fallthrough, so level 1 leaks into the level-2 case and returns
// the wrong label.
package tierbug

// Label returns "low" for 1, "mid" for 2, "high" for 3, else "?".
func Label(level int) string {
	label := "?"
	switch level {
	case 1:
		// CHANGE CODE BELOW THIS LINE
		label = "low"
		fallthrough
		// CHANGE CODE ABOVE THIS LINE
	case 2:
		label = "mid"
	case 3:
		label = "high"
	}
	return label
}
