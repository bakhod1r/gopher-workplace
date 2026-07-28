// Package parsebool parses common truthy/falsey config strings.
// A planted bug misclassifies one accepted form.
package parsebool

import "strings"

// Parse accepts (case-insensitive): true/false, 1/0, yes/no, on/off.
// Returns (value, true) if recognized, else (false, false).
func Parse(s string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "1", "yes", "on":
		return true, true
	// CHANGE CODE BELOW THIS LINE
	case "false", "0", "no":
		// CHANGE CODE ABOVE THIS LINE
		return false, true
	}
	return false, false
}
