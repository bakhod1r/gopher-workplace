// Package readerifc — Gopher Workplace challenge.
package readerifc

// Source yields the next chunk; ok is false once it is drained.
type Source interface {
	Read() (chunk string, ok bool)
}

// StringSource reads a string Chunk bytes at a time.
type StringSource struct {
	Data  string
	Chunk int
	pos   int
}

// Read returns the next chunk.
//
// Examples:
//
//	s := &StringSource{Data: "abc", Chunk: 2}
//	s.Read() => "ab", true
//	s.Read() => "c", true
//	s.Read() => "", false
func (s *StringSource) Read() (string, bool) {
	// TODO(candidate): slice from pos, advance, report exhaustion.
	panic("not implemented")
}

// ReadAll drains the source into one string.
func ReadAll(s Source) string {
	// TODO(candidate): read until ok is false.
	panic("not implemented")
}
