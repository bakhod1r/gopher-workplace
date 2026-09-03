// Package flushwriter — Gopher Workplace challenge.
package flushwriter

import (
	"bufio"
	"io"
)

// WriteAll writes each line followed by '\n' through a buffered writer.
//
// A buffered writer holds the tail of the output until it is flushed; the
// last partial buffer is lost otherwise.
//
// Examples:
//
//	WriteAll(&buf, []string{"a"}) => buf holds "a\n"
func WriteAll(w io.Writer, lines []string) error {
	// CHANGE CODE BELOW THIS LINE
	bw := bufio.NewWriter(w)
	for _, l := range lines {
		if _, err := bw.WriteString(l); err != nil {
			return err
		}
		if err := bw.WriteByte('\n'); err != nil {
			return err
		}
	}
	return nil
	// CHANGE CODE ABOVE THIS LINE
}
