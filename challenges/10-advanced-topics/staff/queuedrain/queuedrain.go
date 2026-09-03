// Package queuedrain — Gopher Workplace challenge.
package queuedrain

import (
	"sync"
	"sync/atomic"
)

// Queue fans work out to a fixed set of workers.
type Queue struct {
	ch    chan int
	wg    sync.WaitGroup
	total atomic.Int64
	once  sync.Once
}

// NewQueue starts n workers, each accumulating the values it receives.
func NewQueue(n int) *Queue {
	if n < 1 {
		n = 1
	}
	q := &Queue{ch: make(chan int, 16)}
	q.wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer q.wg.Done()
			for v := range q.ch {
				q.total.Add(int64(v))
			}
		}()
	}
	return q
}

// Push submits one value. It must not be called after Close.
func (q *Queue) Push(v int) { q.ch <- v }

// Close stops accepting work, waits for the workers to finish what is
// already queued, and returns the total they processed.
//
// Every goroutine the Queue started must have exited by the time Close
// returns; nothing may be left blocked on the channel.
//
// Examples:
//
//	q := NewQueue(4); q.Push(1); q.Close() => 1
func (q *Queue) Close() int64 {
	panic("not implemented")
}
