// Package latencypeak - Gopher Workplace challenge.
package latencypeak

import "sync/atomic"

// PeakTracker records the slowest request latency observed, in milliseconds.
type PeakTracker struct {
	ms atomic.Int64
}

// Observe raises the tracked peak to ms when ms is larger.
//
// Examples:
//
//	var p PeakTracker; p.Observe(5); p.Peak()               => 5
//	var p PeakTracker; p.Observe(5); p.Observe(3); p.Peak() => 5
func (p *PeakTracker) Observe(ms int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Peak returns the slowest observed latency, or 0 if none.
//
// Examples:
//
//	var p PeakTracker; p.Peak() => 0
func (p *PeakTracker) Peak() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}
