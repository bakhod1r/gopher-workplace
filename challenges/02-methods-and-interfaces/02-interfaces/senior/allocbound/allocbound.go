// Package allocbound — Gopher Workplace challenge.
package allocbound

// Recorder accepts records.
type Recorder interface {
	Write(id int)
}

// Sink stores ids in a preallocated buffer.
type Sink struct {
	buf []int
}

// NewSink returns a sink with capacity reserved for n records.
func NewSink(n int) *Sink {
	// TODO(candidate): reserve capacity up front.
	panic("not implemented")
}

// Write appends one id.
func (s *Sink) Write(id int) {
	// TODO(candidate): append without regrowing when capacity remains.
	panic("not implemented")
}

// Len returns how many records were written.
func (s *Sink) Len() int {
	// TODO(candidate): number of records.
	panic("not implemented")
}

// Fill writes ids 0..n-1 through the recorder.
//
// Examples:
//
//	s := NewSink(1000); Fill(s, 1000) => s.Len() == 1000
func Fill(r Recorder, n int) {
	// TODO(candidate): n writes, no extra allocation per write.
	panic("not implemented")
}
