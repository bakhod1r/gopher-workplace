# Bounded Blocking Queue

## Intuition

A condition variable is the primitive for "sleep until this predicate becomes true". Wakeups are hints, not guarantees, so the predicate always lives in a `for` loop — and closing has to broadcast, because every waiter needs to re-check.

## Approach

1. `Push` waits while full and not closed, then rechecks `closed` before enqueuing.
2. `Pop` waits while empty and not closed, then returns false if still empty.
3. Each operation signals the opposite condition after changing the state.
4. `Close` sets the flag and broadcasts both conditions under the lock.

## Solution

```go
func (q *Queue) Push(v int) bool {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) >= q.cap && !q.closed {
		q.notFull.Wait()
	}
	if q.closed {
		return false
	}

	q.items = append(q.items, v)
	q.notEmpty.Signal()
	return true
}

func (q *Queue) Pop() (int, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	for len(q.items) == 0 && !q.closed {
		q.notEmpty.Wait()
	}
	if len(q.items) == 0 {
		return 0, false
	}

	v := q.items[0]
	q.items = q.items[1:]
	q.notFull.Signal()
	return v, true
}

func (q *Queue) Close() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.closed = true
	q.notFull.Broadcast()
	q.notEmpty.Broadcast()
}
```

## Walkthrough

`Pop` checks emptiness rather than `closed` after waking, which is what lets a closed queue still drain its backlog before reporting false.

## Pitfalls

- `if` instead of `for` around `Wait` — a spurious or stale wakeup proceeds with a false predicate.
- `Signal` in `Close` instead of `Broadcast`, leaving all but one waiter asleep forever.
- Calling `Wait` without holding the mutex, which panics.
