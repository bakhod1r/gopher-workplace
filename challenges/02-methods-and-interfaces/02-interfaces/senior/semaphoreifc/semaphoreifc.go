// Package semaphoreifc — Gopher Workplace challenge.
package semaphoreifc

import "sync"

// Limiter bounds concurrent access.
type Limiter interface {
	Acquire()
	Release()
}

// Semaphore permits at most N concurrent holders.
type Semaphore struct {
	slots chan struct{}
}

// NewSemaphore returns a semaphore with n slots.
func NewSemaphore(n int) *Semaphore {
	return &Semaphore{slots: make(chan struct{}, n)}
}

// Acquire takes a slot, blocking until one is free.
func (s *Semaphore) Acquire() {
	// TODO(candidate): take a slot.
	panic("not implemented")
}

// TryAcquire takes a slot if one is free, without blocking.
func (s *Semaphore) TryAcquire() bool {
	// TODO(candidate): non-blocking acquire.
	panic("not implemented")
}

// Release returns a slot.
func (s *Semaphore) Release() {
	// TODO(candidate): return a slot.
	panic("not implemented")
}

// Job is one unit of work.
type Job interface {
	Do()
}

// RunLimited runs every job concurrently, at most limit at a time.
//
// Examples:
//
//	limit 2 over 100 jobs => peak concurrency 2
func RunLimited(jobs []Job, limit int) {
	// TODO(candidate): semaphore-bounded fan-out; wait for all jobs.
	panic("not implemented")
}

var _ = sync.WaitGroup{}
