// Package shadowerr parses a decimal string into an int with a named error
// return. A planted bug shadows the named err inside the if block, so a real
// parse failure is reported as success.
package shadowerr

import "strconv"

// Parse returns the parsed number and an error. On failure n is 0 and err is
// non-nil.
func Parse(s string) (n int, err error) {
	// CHANGE CODE BELOW THIS LINE
	if v, err := strconv.Atoi(s); err == nil {
		n = v
	}
	// CHANGE CODE ABOVE THIS LINE
	return
}
