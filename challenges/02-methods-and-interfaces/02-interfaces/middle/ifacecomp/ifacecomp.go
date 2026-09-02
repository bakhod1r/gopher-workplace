// Package ifacecomp — Gopher Workplace challenge.
package ifacecomp

// Reader yields stored text.
type Reader interface {
	Read() string
}

// Writer stores text.
type Writer interface {
	Write(s string)
}

// ReadWriter does both.
type ReadWriter interface {
	Reader
	Writer
}

// File can read and write.
type File struct {
	data string
}

// Read returns the stored text.
func (f *File) Read() string {
	// TODO(candidate): return the data.
	panic("not implemented")
}

// Write replaces the stored text.
func (f *File) Write(s string) {
	// TODO(candidate): store the data.
	panic("not implemented")
}

// ReadOnly can only read.
type ReadOnly struct {
	Data string
}

// Read returns the stored text.
func (r ReadOnly) Read() string { return r.Data }

// WriteOnly can only write.
type WriteOnly struct {
	Sink []string
}

// Write records the text.
func (w *WriteOnly) Write(s string) { w.Sink = append(w.Sink, s) }

// Describe reports which capabilities v has: "rw", "r", "w" or "none".
//
// Examples:
//
//	Describe(&File{})              => "rw"
//	Describe(ReadOnly{Data: "x"})  => "r"
func Describe(v any) string {
	// TODO(candidate): assert to Reader and Writer, then combine.
	panic("not implemented")
}
