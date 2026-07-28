// Package utf8count counts runes by scanning lead bytes.
// A planted bug counts continuation bytes instead.
package utf8count

// Count returns the number of runes in a valid UTF-8 byte slice by counting
// every byte that is NOT a continuation byte (10xxxxxx).
func Count(b []byte) int {
	n := 0
	for _, c := range b {
		// CHANGE CODE BELOW THIS LINE
		if c&0xC0 == 0x80 {
			// CHANGE CODE ABOVE THIS LINE
			n++
		}
	}
	return n
}
