// Package maskcard masks all but the last four characters.
// A planted bug reveals the wrong digits.
package maskcard

// Mask replaces every character except the last four with '*'. If s has four or
// fewer characters, it is returned unchanged (nothing to hide).
func Mask(s string) string {
	r := []rune(s)
	if len(r) <= 4 {
		return s
	}
	out := make([]rune, len(r))
	for i := range r {
		// CHANGE CODE BELOW THIS LINE
		if i >= len(r)-4 {
			// CHANGE CODE ABOVE THIS LINE
			out[i] = '*'
		} else {
			out[i] = r[i]
		}
	}
	return string(out)
}
