// Package writerifc — Gopher Workplace challenge.
package writerifc

// Writer accepts text and reports how many bytes it took.
type Writer interface {
	Write(s string) int
}

// Builder collects everything written to it.
type Builder struct {
	buf string
}

// Write appends s and returns its byte length.
//
// Examples:
//
//	b := &Builder{}; b.Write("ab") => 2
func (b *Builder) Write(s string) int {
	// TODO(candidate): append, return the byte count.
	panic("not implemented")
}

// String returns everything written so far.
func (b *Builder) String() string {
	// TODO(candidate): return the buffer.
	panic("not implemented")
}

// WriteLines writes each line plus a newline and returns the total bytes.
func WriteLines(w Writer, lines []string) int {
	// TODO(candidate): write line + "\n", summing the counts.
	panic("not implemented")
}
