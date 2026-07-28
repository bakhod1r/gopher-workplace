// Package wordwrap breaks text into lines no longer than width runes.
// A planted off-by-one lets a line exceed the width.
package wordwrap

import "strings"

// Wrap joins words with spaces into lines whose length is at most width runes,
// breaking before a word that would overflow. Assumes single-space input, no
// word longer than width.
func Wrap(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	var line string
	for _, w := range words {
		if line == "" {
			line = w
			continue
		}
		// CHANGE CODE BELOW THIS LINE
		if len(line)+1+len(w) < width {
			// CHANGE CODE ABOVE THIS LINE
			line += " " + w
		} else {
			lines = append(lines, line)
			line = w
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
