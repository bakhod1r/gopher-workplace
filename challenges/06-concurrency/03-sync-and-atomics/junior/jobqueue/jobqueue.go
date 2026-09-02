// Package jobqueue - Gopher Workplace challenge.
package jobqueue

import "sync"

// JobQueue is a blocking FIFO queue feeding a worker pool.
type JobQueue struct {
	mu       sync.Mutex
	notEmpty *sync.Cond
	jobs     []string
	closed   bool
}

// NewJobQueue returns an open, empty queue.
func NewJobQueue() *JobQueue {
	q := &JobQueue{}
	q.notEmpty = sync.NewCond(&q.mu)
	return q
}

// Submit enqueues a job and wakes one waiting worker.
//
// Examples:
//
//	q := NewJobQueue(); q.Submit("a"); q.Take() => "a", true
func (q *JobQueue) Submit(job string) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Take returns the oldest job, blocking while the queue is empty and open.
// It reports false once the queue is closed and drained.
//
// Examples:
//
//	q.Submit("a"); q.Submit("b"); q.Take() => "a", true
//	q := NewJobQueue(); q.Close(); q.Take() => "", false
func (q *JobQueue) Take() (string, bool) {
	// TODO(candidate): implement this.
	panic("not implemented")
}

// Close closes the queue and wakes every waiting worker.
func (q *JobQueue) Close() {
	// TODO(candidate): implement this.
	panic("not implemented")
}
