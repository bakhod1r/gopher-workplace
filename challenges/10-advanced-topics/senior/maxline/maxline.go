// Package maxline — Gopher Workplace challenge.
package maxline

import "io"

// MaxLine returns the length in bytes of the longest '\n'-separated line
// in r, not counting the newline itself.
//
// Lines may be longer than any single read, and the stream may be far
// larger than memory. Only a fixed-size buffer may be held.
//
// Examples:
//
//	MaxLine(strings.NewReader("ab\ncdef\n")) => 4, nil
func MaxLine(r io.Reader) (int, error) {
	panic("not implemented")
}
