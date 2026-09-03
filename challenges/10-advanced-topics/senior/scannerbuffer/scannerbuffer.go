// Package scannerbuffer — Gopher Workplace challenge.
package scannerbuffer

import (
	"bufio"
	"io"
)

// maxLine is the longest line this reader must accept.
const maxLine = 4 << 20

// LongestLine returns the length of the longest line in r.
//
// bufio.Scanner refuses tokens larger than its buffer limit, which defaults
// to 64 KiB. A line longer than that is an error, not a truncation.
//
// Examples:
//
//	LongestLine(strings.NewReader("ab\ncdef")) => 4, nil
func LongestLine(r io.Reader) (int, error) {
	// CHANGE CODE BELOW THIS LINE
	sc := bufio.NewScanner(r)
	best := 0
	for sc.Scan() {
		if n := len(sc.Bytes()); n > best {
			best = n
		}
	}
	if err := sc.Err(); err != nil {
		return 0, err
	}
	return best, nil
	// CHANGE CODE ABOVE THIS LINE
}
