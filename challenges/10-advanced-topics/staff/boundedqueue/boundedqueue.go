// Package boundedqueue — Gopher Workplace challenge.
package boundedqueue

// Queue is a bounded FIFO with blocking producers.
type Queue struct {
	ch chan int
}

// NewQueue returns a queue holding at most n values.
func NewQueue(n int) *Queue {
	if n < 1 {
		n = 1
	}
	return &Queue{ch: make(chan int, n)}
}

// Take removes and returns the oldest value, waiting if the queue is empty.
func (q *Queue) Take(done <-chan struct{}) (int, bool) {
	select {
	case v := <-q.ch:
		return v, true
	case <-done:
		return 0, false
	}
}

// Len reports how many values are queued.
func (q *Queue) Len() int { return len(q.ch) }

// Cap reports the queue's capacity.
func (q *Queue) Cap() int { return cap(q.ch) }

// Put appends v to the queue, waiting while the queue is full, and
// reports whether it was accepted.
//
// Waiting is what applies backpressure; the queue must never grow past its
// capacity, and a cancelled producer must not wait forever.
//
// Examples:
//
//	q := NewQueue(2); q.Put(done, 1) => true
func (q *Queue) Put(done <-chan struct{}, v int) bool {
	panic("not implemented")
}
