// Package chunkreader — Gopher Workplace challenge.
package chunkreader

// Reader fills the caller's buffer and reports how many bytes it wrote.
// A return of (0, false) means the source is drained.
type Reader interface {
	Read(p []byte) (int, bool)
}

// ChunkSource serves a string through a caller-supplied buffer.
type ChunkSource struct {
	Data string
	pos  int
}

// Read copies the next bytes into p.
//
// Examples:
//
//	s := &ChunkSource{Data: "abcd"}
//	s.Read(make([]byte, 3)) => 3, true   // "abc"
//	s.Read(make([]byte, 3)) => 1, true   // "d"
//	s.Read(make([]byte, 3)) => 0, false
func (s *ChunkSource) Read(p []byte) (int, bool) {
	// TODO(candidate): copy into p, advance, report the count.
	panic("not implemented")
}

// CountLines counts newline bytes, streaming through buf.
//
// buf is reused for every chunk; CountLines must not allocate per chunk.
func CountLines(r Reader, buf []byte) int {
	// TODO(candidate): read chunks into buf, count '\n'.
	panic("not implemented")
}
