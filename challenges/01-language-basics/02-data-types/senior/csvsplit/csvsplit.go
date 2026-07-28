// Package csvsplit splits one CSV line into fields, honoring double quotes.
// A planted bug splits on a comma that is inside a quoted field.
package csvsplit

// Split parses a single RFC-4180-ish line: fields separated by commas, a field
// may be wrapped in double quotes to contain commas. Quotes are not included in
// the output (inner "" -> ").
func Split(line string) []string {
	var fields []string
	var cur []byte
	inQuotes := false
	for i := 0; i < len(line); i++ {
		c := line[i]
		switch {
		case c == '"':
			if inQuotes && i+1 < len(line) && line[i+1] == '"' {
				cur = append(cur, '"')
				i++
			} else {
				inQuotes = !inQuotes
			}
		// CHANGE CODE BELOW THIS LINE
		case c == ',':
			// CHANGE CODE ABOVE THIS LINE
			fields = append(fields, string(cur))
			cur = cur[:0]
		default:
			cur = append(cur, c)
		}
	}
	fields = append(fields, string(cur))
	return fields
}
