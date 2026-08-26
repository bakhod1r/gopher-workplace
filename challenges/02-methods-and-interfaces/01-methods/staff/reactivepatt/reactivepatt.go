// Package reactivepatt — Gopher Workplace challenge.
package reactivepatt

// Stream represents a stream of ints.
type Stream struct {
	Data []int
}

// Filter modifies the stream in place, keeping only elements satisfying fn.
func (s *Stream) Filter(fn func(int) bool) *Stream {
	// TODO(candidate): filter s.Data in place, return s
	panic("not implemented")
}

// Map modifies the stream in place, applying fn to each.
func (s *Stream) Map(fn func(int) int) *Stream {
	// TODO(candidate): map s.Data in place, return s
	panic("not implemented")
}
