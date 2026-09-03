// Package bufferpool — Gopher Workplace challenge.
package bufferpool

import "sync"

// Encoder renders log lines, reusing scratch buffers from a pool so a busy
// shipper does not allocate one per line.
type Encoder struct {
	pool sync.Pool
}

// NewEncoder returns an Encoder whose pool hands out empty byte slices.
//
// Examples:
//
//	NewEncoder() != nil => true
func NewEncoder() *Encoder {
	// TODO(candidate): build the pool with a New func that returns an empty
	// []byte with some spare capacity.
	panic("not implemented")
}

// Encode joins fields with '|' using a buffer taken from the pool, then
// returns the buffer to the pool.
//
// Examples:
//
//	NewEncoder().Encode([]string{"warn", "disk full"}) => "warn|disk full"
//	NewEncoder().Encode([]string{"solo"})              => "solo"
//	NewEncoder().Encode(nil)                           => ""
func (e *Encoder) Encode(fields []string) string {
	// TODO(candidate): take a buffer, reset its length, append the fields,
	// build the string, put the buffer back.
	panic("not implemented")
}
