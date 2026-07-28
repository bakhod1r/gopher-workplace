// Package splitlines splits text into lines, handling both \n and \r\n.
// A planted bug leaves a trailing carriage return on CRLF input.
package splitlines

import "strings"

// Lines splits s on '\n' and strips a trailing '\r' from each line so that
// Windows (CRLF) and Unix (LF) input yield identical results.
func Lines(s string) []string {
	parts := strings.Split(s, "\n")
	for i, p := range parts {
		// CHANGE CODE BELOW THIS LINE
		parts[i] = p
		// CHANGE CODE ABOVE THIS LINE
	}
	return parts
}
