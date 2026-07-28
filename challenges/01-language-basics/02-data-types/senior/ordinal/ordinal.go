// Package ordinal formats an integer as an English ordinal (1st, 2nd, 3rd...).
// A planted bug mishandles the 11-13 exception.
package ordinal

import "strconv"

// Suffix returns the ordinal suffix ("st","nd","rd","th") for a non-negative n.
func Suffix(n int) string {
	// CHANGE CODE BELOW THIS LINE
	switch n % 10 {
	// CHANGE CODE ABOVE THIS LINE
	case 1:
		return "st"
	case 2:
		return "nd"
	case 3:
		return "rd"
	default:
		return "th"
	}
}

// Format returns n with its ordinal suffix, e.g. "21st".
func Format(n int) string { return strconv.Itoa(n) + Suffix(n) }
