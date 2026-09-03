// Package streamcount — Gopher Workplace challenge.
package streamcount

import (
	"bytes"
	"io"
)

// CountLines returns the number of '\n' bytes in r.
//
// The reader may deliver gigabytes. The function must work in one pass over
// a fixed-size buffer and must never hold the whole stream in memory.
//
// Examples:
//
//	CountLines(strings.NewReader("a\nb\n")) => 2, nil
func CountLines(r io.Reader) (int, error) {
	panic("not implemented")
}
