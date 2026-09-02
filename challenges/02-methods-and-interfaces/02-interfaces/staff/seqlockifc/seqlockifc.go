// Package seqlockifc — Gopher Workplace challenge.
package seqlockifc

import (
	"sync"
	"sync/atomic"
)

// Snapshot is a pair of related counters that must be read consistently.
type Snapshot struct {
	Requests int64
	Errors   int64
}

// Reader returns a consistent snapshot.
type Reader interface {
	Read() Snapshot
}

// SeqLock protects a snapshot with a sequence counter.
type SeqLock struct {
	seq atomic.Uint64 // even: stable, odd: write in progress

	mu   sync.Mutex // serialises writers
	reqs atomic.Int64
	errs atomic.Int64
}

// Write updates both fields as one logical update.
//
// Examples:
//
//	Write(10, 2) => Read() returns {10, 2}
func (s *SeqLock) Write(requests, errors int64) {
	// TODO(candidate): seq to odd, update, seq to even.
	panic("not implemented")
}

// Read returns a consistent snapshot, retrying while a write is in progress.
func (s *SeqLock) Read() Snapshot {
	// TODO(candidate): optimistic read validated by the sequence.
	panic("not implemented")
}

// Seq returns the current sequence value.
func (s *SeqLock) Seq() uint64 {
	// TODO(candidate): current sequence.
	panic("not implemented")
}
