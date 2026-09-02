// Package ringsink — Gopher Workplace challenge.
package ringsink

// Sink accepts log lines.
type Sink interface {
	Write(line string)
}

// RingSink keeps the last N lines in a fixed-size ring.
type RingSink struct {
	buf   []string
	count int // total writes ever
}

// NewRingSink returns a ring holding the last size lines.
func NewRingSink(size int) *RingSink {
	if size < 1 {
		size = 1
	}
	return &RingSink{buf: make([]string, size)}
}

// Write stores a line, overwriting the oldest once full.
//
// Examples:
//
//	size 3; write a, b, c, d => Snapshot is [b c d]
func (r *RingSink) Write(line string) {
	// TODO(candidate): store at count%len, then advance count.
	panic("not implemented")
}

// Len returns how many lines are currently held.
func (r *RingSink) Len() int {
	// TODO(candidate): min(count, size).
	panic("not implemented")
}

// Snapshot returns the held lines, oldest first.
func (r *RingSink) Snapshot() []string {
	// TODO(candidate): walk from the oldest index forward.
	panic("not implemented")
}
