// Package sinkfanout — Gopher Workplace challenge.
package sinkfanout

import (
	"errors"
	"sync"
)

// ErrSinkFailed is returned by a broken sink.
var ErrSinkFailed = errors.New("sink failed")

// Sink accepts events.
type Sink interface {
	Write(event string) error
}

// MemSink records events; it is safe for concurrent use.
type MemSink struct {
	mu     sync.Mutex
	Events []string
}

// Write records the event.
func (m *MemSink) Write(event string) error {
	// TODO(candidate): append under the mutex.
	panic("not implemented")
}

// Len returns how many events were recorded.
func (m *MemSink) Len() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.Events)
}

// ErrSink always fails.
type ErrSink struct{}

// Write always returns ErrSinkFailed.
func (ErrSink) Write(event string) error {
	// TODO(candidate): always fail.
	panic("not implemented")
}

// FanOut writes the event to every sink concurrently and returns the number
// of sinks that failed.
//
// Examples:
//
//	two good sinks and one failing => 1
func FanOut(sinks []Sink, event string) int {
	// TODO(candidate): concurrent fan-out, race-free failure count.
	panic("not implemented")
}
