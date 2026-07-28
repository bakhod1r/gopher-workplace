// Package hexdecode decodes a lowercase/uppercase hex string to bytes.
// A planted bug swaps the nibble order.
package hexdecode

func nib(c byte) (int, bool) {
	switch {
	case c >= '0' && c <= '9':
		return int(c - '0'), true
	case c >= 'a' && c <= 'f':
		return int(c-'a') + 10, true
	case c >= 'A' && c <= 'F':
		return int(c-'A') + 10, true
	}
	return 0, false
}

// Decode parses hex (even length) into bytes. Returns (nil,false) on bad input.
func Decode(s string) ([]byte, bool) {
	if len(s)%2 != 0 {
		return nil, false
	}
	out := make([]byte, 0, len(s)/2)
	for i := 0; i < len(s); i += 2 {
		hi, ok1 := nib(s[i])
		lo, ok2 := nib(s[i+1])
		if !ok1 || !ok2 {
			return nil, false
		}
		// CHANGE CODE BELOW THIS LINE
		out = append(out, byte(lo<<4|hi))
		// CHANGE CODE ABOVE THIS LINE
	}
	return out, true
}
