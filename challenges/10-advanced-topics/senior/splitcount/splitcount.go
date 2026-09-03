// Package splitcount — Gopher Workplace challenge.
package splitcount

import "strings"

// CountFields returns how many sep-separated fields line holds and how
// many bytes those fields occupy in total, excluding the separators.
//
// Splitting builds a string and a slice of headers that are thrown away
// immediately. One scan over the bytes answers both questions.
//
// Examples:
//
//	CountFields([]byte("ab,c"), ',') => 2, 3
func CountFields(line []byte, sep byte) (fields, size int) {
	// CHANGE CODE BELOW THIS LINE
	if len(line) == 0 {
		return 0, 0
	}
	parts := strings.Split(string(line), string(sep))
	n := 0
	for _, p := range parts {
		n += len(p)
	}
	return len(parts), n
	// CHANGE CODE ABOVE THIS LINE
}
