// Package chunkedreader — Gopher Workplace challenge.
package chunkedreader

import "io"

// ReadAll drains r into dst, reusing dst's capacity and growing it only when
// needed, and returns the filled slice. Reads happen in chunks of chunk
// bytes; a non-positive chunk uses 4096. io.EOF ends the read normally, any
// other error is returned along with what was read so far.
//
// Examples:
//
//	ReadAll(strings.NewReader("hello"), nil, 2) => []byte("hello"), nil
func ReadAll(r io.Reader, dst []byte, chunk int) ([]byte, error) {
	panic("not implemented")
}

// CountChunks reports how many Read calls draining n bytes in chunks of size
// chunk takes, including the final call that reports EOF. A non-positive
// chunk or n gives 0.
//
// Examples:
//
//	CountChunks(5, 2) => 4
func CountChunks(n, chunk int) int {
	panic("not implemented")
}
