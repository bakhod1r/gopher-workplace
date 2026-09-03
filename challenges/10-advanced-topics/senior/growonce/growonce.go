// Package growonce — Gopher Workplace challenge.
package growonce

import "io"

// Collect reads r to EOF and returns its bytes.
//
// hint is the caller's estimate of the size. When it is accurate the whole
// read must cost a single allocation instead of a chain of doublings.
//
// Examples:
//
//	Collect(strings.NewReader("abc"), 3) => []byte("abc"), nil
func Collect(r io.Reader, hint int) ([]byte, error) {
	panic("not implemented")
}
