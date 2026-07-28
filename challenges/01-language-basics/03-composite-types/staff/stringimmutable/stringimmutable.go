// Package stringimmutable upper-cases ASCII. A planted bug mutates a []byte copy
// but returns the original string, which is immutable and unchanged.
package stringimmutable

// Upper returns s with ASCII lowercase letters upper-cased.
func Upper(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'a' && c <= 'z' {
			b[i] = c - 32
		}
	}
	// CHANGE CODE BELOW THIS LINE
	return s
	// CHANGE CODE ABOVE THIS LINE
}
