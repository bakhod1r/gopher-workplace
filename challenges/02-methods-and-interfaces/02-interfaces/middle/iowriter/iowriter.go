// Package iowriter — Gopher Workplace challenge.
package iowriter

import "io"

// WriteReport writes a title line followed by one line per item.
//
// The output is "<title>\n" then "- <item>\n" for each item.
// It returns the number of bytes written and the first error, if any.
//
// Examples:
//
//	WriteReport(buf, "T", []string{"a"}) => 6, nil   // "T\n- a\n"
//	WriteReport(buf, "T", nil)           => 2, nil   // "T\n"
func WriteReport(w io.Writer, title string, items []string) (int, error) {
	// TODO(candidate): write incrementally, summing byte counts.
	panic("not implemented")
}
