// Package utf8valid validates a UTF-8 byte sequence by hand.
// A planted bug in the 2-byte lead mask misclassifies longer sequences.
package utf8valid

// Valid reports whether b is well-formed UTF-8 (correct lead bytes and the right
// number of 10xxxxxx continuation bytes). Overlong forms are not specifically
// rejected here.
func Valid(b []byte) bool {
	i := 0
	for i < len(b) {
		c := b[i]
		var n int
		switch {
		case c < 0x80:
			n = 0
		// CHANGE CODE BELOW THIS LINE
		case c&0xC0 == 0xC0:
			// CHANGE CODE ABOVE THIS LINE
			n = 1
		case c&0xF0 == 0xE0:
			n = 2
		case c&0xF8 == 0xF0:
			n = 3
		default:
			return false
		}
		for j := 1; j <= n; j++ {
			if i+j >= len(b) || b[i+j]&0xC0 != 0x80 {
				return false
			}
		}
		i += n + 1
	}
	return true
}
