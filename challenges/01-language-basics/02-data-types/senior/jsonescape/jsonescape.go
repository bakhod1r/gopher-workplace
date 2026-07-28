// Package jsonescape escapes a string for embedding in a JSON document.
// A planted bug forgets to escape the backslash itself.
package jsonescape

import "strings"

// Escape returns s with JSON-required escapes: backslash, double-quote, newline,
// tab, and carriage return. (Other control chars are ignored for simplicity.)
func Escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		// CHANGE CODE BELOW THIS LINE
		case '"':
			b.WriteString(`\"`)
		// CHANGE CODE ABOVE THIS LINE
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\t`)
		case '\r':
			b.WriteString(`\r`)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
