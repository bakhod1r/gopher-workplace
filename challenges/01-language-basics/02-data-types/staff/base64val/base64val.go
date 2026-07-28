// Package base64val maps a standard base64 character to its 6-bit value.
// A planted bug offsets the digit range by one.
package base64val

// Value returns the 0..63 value of a base64 character (A-Z, a-z, 0-9, +, /),
// or (0,false) for any other byte.
func Value(c byte) (int, bool) {
	switch {
	case c >= 'A' && c <= 'Z':
		return int(c - 'A'), true
	case c >= 'a' && c <= 'z':
		return int(c-'a') + 26, true
	// CHANGE CODE BELOW THIS LINE
	case c >= '0' && c <= '9':
		return int(c-'0') + 53, true
	// CHANGE CODE ABOVE THIS LINE
	case c == '+':
		return 62, true
	case c == '/':
		return 63, true
	}
	return 0, false
}
