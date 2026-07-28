// Package tolowerascii lowercases ASCII letters only.
// A planted bug also alters characters just past 'Z'.
package tolowerascii

// Lower returns s with A-Z mapped to a-z; every other byte is unchanged.
func Lower(s string) string {
	b := []byte(s)
	for i, c := range b {
		// CHANGE CODE BELOW THIS LINE
		if c >= 'A' {
			// CHANGE CODE ABOVE THIS LINE
			b[i] = c + 32
		}
	}
	return string(b)
}
