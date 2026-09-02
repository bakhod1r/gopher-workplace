// Package boundedq — Gopher Workplace challenge.
package boundedq

import "sync"

// Queue is a bounded blocking FIFO queue.
type Queue struct {
	mu       sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond

	items  []int
	cap    int
	closed bool
}

// Sink accepts items.
type Sink interface {
	Push(v int) bool
}

// NewQueue returns a queue holding at most capacity items.
func NewQueue(capacity int) *Queue {
	if capacity < 1 {
		capacity = 1
	}
	q := &Queue{cap: capacity}
	q.notFull = sync.NewCond(&q.mu)
	q.notEmpty = sync.NewCond(&q.mu)
	return q
}

// Push adds an item, blocking while the queue is full.
// It returns false when the queue is closed.
//
// Examples:
//
//	capacity 1; Push(1) => true
func (q *Queue) Push(v int) bool {
	// TODO(candidate): wait for room, then enqueue and signal.
	panic("not implemented")
}

// Pop removes an item, blocking while the queue is empty.
// It returns false once the queue is closed and drained.
func (q *Queue) Pop() (int, bool) {
	// TODO(candidate): wait for an item, then dequeue and signal.
	panic("not implemented")
}

// Close closes the queue and wakes every blocked goroutine.
func (q *Queue) Close() {
	// TODO(candidate): mark closed, broadcast both conditions.
	panic("not implemented")
}

// Len returns the number of queued items.
func (q *Queue) Len() int {
	// TODO(candidate): read under the lock.
	panic("not implemented")
}
