// Package sliceleak — Gopher Workplace challenge.
package sliceleak

// Extractor pulls a prefix out of a buffer.
type Extractor interface {
	Extract(b []byte, n int) []byte
}

// Aliasing keeps a view into the source buffer.
type Aliasing struct{}

// Extract returns a sub-slice of b.
func (Aliasing) Extract(b []byte, n int) []byte { return Prefix(b, n) }

// Copying detaches from the source buffer.
type Copying struct{}

// Extract returns an independent copy.
func (Copying) Extract(b []byte, n int) []byte { return PrefixCopy(b, n) }

// Prefix returns the first n bytes as a sub-slice.
//
// The result aliases b and keeps its whole backing array alive.
//
// Examples:
//
//	Prefix(make([]byte, 1<<20), 8) => len 8, cap 1<<20
func Prefix(b []byte, n int) []byte {
	// TODO(candidate): clamp n, then sub-slice.
	panic("not implemented")
}

// PrefixCopy returns the first n bytes as an independent copy.
func PrefixCopy(b []byte, n int) []byte {
	// TODO(candidate): clamp n, then copy exactly n bytes.
	panic("not implemented")
}

// RetainedBytes reports how much memory the result still pins.
func RetainedBytes(b []byte) int {
	// TODO(candidate): the backing array size the slice keeps reachable.
	panic("not implemented")
}
