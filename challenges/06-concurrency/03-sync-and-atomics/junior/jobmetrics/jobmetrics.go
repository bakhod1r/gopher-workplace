// Package jobmetrics - Gopher Workplace challenge.
package jobmetrics

import "sync/atomic"

// JobMetrics is a lock-free counter of jobs processed by a worker pool.
type JobMetrics struct {
	processed atomic.Int64
}

// Add adds delta to the processed-jobs counter.
//
// Examples:
//
//	var m JobMetrics; m.Add(3); m.Processed()            => 3
//	var m JobMetrics; m.Add(2); m.Add(-5); m.Processed() => -3
func (m *JobMetrics) Add(delta int64) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Processed returns the current number of processed jobs.
func (m *JobMetrics) Processed() int64 {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Reset zeroes the counter after a metrics flush.
//
// Examples:
//
//	var m JobMetrics; m.Add(7); m.Reset(); m.Processed() => 0
func (m *JobMetrics) Reset() {
	// TODO(candidate): implement this.
	panic("not implemented")
}
