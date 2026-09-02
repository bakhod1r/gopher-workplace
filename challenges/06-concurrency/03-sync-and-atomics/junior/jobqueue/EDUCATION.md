# Worker Job Queue

## Intuition

`Wait` must be called with the lock held: it releases the lock while parked so a producer can make progress, and re-acquires it before returning. Because another woken worker may have taken the job first, the condition has to be re-checked in a loop.

## Approach

1. `Submit` locks, appends, unlocks, then `Signal`s one worker.
2. `Take` locks, waits in a `for` loop while empty and open, then pops the front job.
3. `Close` locks, sets `closed`, unlocks, then `Broadcast`s so every worker exits.

## Solution

```go
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
	q.mu.Lock()
	q.jobs = append(q.jobs, job)
	q.mu.Unlock()
	q.notEmpty.Signal()
}

// Take returns the oldest job, blocking while the queue is empty and open.
// It reports false once the queue is closed and drained.
//
// Examples:
//
//	q.Submit("a"); q.Submit("b"); q.Take() => "a", true
//	q := NewJobQueue(); q.Close(); q.Take() => "", false
func (q *JobQueue) Take() (string, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()
	for len(q.jobs) == 0 && !q.closed {
		q.notEmpty.Wait()
	}
	if len(q.jobs) == 0 {
		return "", false
	}
	job := q.jobs[0]
	q.jobs = q.jobs[1:]
	return job, true
}

// Close closes the queue and wakes every waiting worker.
func (q *JobQueue) Close() {
	q.mu.Lock()
	q.closed = true
	q.mu.Unlock()
	q.notEmpty.Broadcast()
}
```

## Walkthrough

Three workers call `Take` on an empty queue and all park inside `Wait`. `Submit("a")` appends and signals: one worker wakes, re-checks the loop condition, sees one job, and pops it. The other two stay parked until `Close` broadcasts.

## Pitfalls

- Using `if` instead of `for` around `Wait` — the job may be gone by the time you wake.
- Calling `Wait` without holding the mutex; that panics.
- Using `Signal` in `Close`, which leaves all but one worker parked forever.
