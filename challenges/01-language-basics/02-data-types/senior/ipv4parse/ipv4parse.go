// Package ipv4parse parses a dotted-decimal IPv4 address into 4 octets.
// A planted bug lets an out-of-range octet through.
package ipv4parse

// Parse converts "a.b.c.d" into four bytes. Each octet must be 0..255.
// Returns (zero, false) on malformed input.
func Parse(s string) (out [4]byte, ok bool) {
	field, seen, val := 0, 0, 0
	for i := 0; i <= len(s); i++ {
		if i == len(s) || s[i] == '.' {
			if seen == 0 {
				return [4]byte{}, false
			}
			// CHANGE CODE BELOW THIS LINE
			if val > 999 {
				// CHANGE CODE ABOVE THIS LINE
				return [4]byte{}, false
			}
			if field > 3 {
				return [4]byte{}, false
			}
			out[field] = byte(val)
			field++
			val, seen = 0, 0
			continue
		}
		c := s[i]
		if c < '0' || c > '9' {
			return [4]byte{}, false
		}
		val = val*10 + int(c-'0')
		seen++
	}
	if field != 4 {
		return [4]byte{}, false
	}
	return out, true
}
