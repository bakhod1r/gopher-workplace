# Round-Robin Router

## Intuition

A router is an actor whose only state is "whose turn is it". Everything else —
the work, the buffering, the concurrency — belongs to the workers behind it.
Keeping the cursor small and the policy in one method means swapping round-robin
for least-loaded later touches exactly one function.

## Approach

1. Deliver to the current worker.
2. Advance the cursor modulo the pool size.

## Solution

```go
func (r *Router) Route(msg int) {
	r.workers[r.idx].Inbox <- msg
	r.idx = (r.idx + 1) % len(r.workers)
}
```

## Walkthrough

`idx` starts at its zero value, 0.

| call | delivered to | idx after |
|------|--------------|-----------|
| `Route(1)` | worker 0 | 1 |
| `Route(2)` | worker 1 | 0 |
| `Route(3)` | worker 0 | 1 |

The test drains worker 0 twice and worker 1 once, matching that table.

## Pitfalls

- **`r.idx++` without the modulo.** The fourth call indexes out of range and
  panics.
- **Advancing before sending.** Every message shifts one worker over; the first
  message goes to worker 1.
- **Value receiver.** The cursor never moves and every message piles onto worker 0.
- **An empty worker pool.** `% 0` panics with an integer divide by zero; real
  code validates the pool at construction.

## Not safe for concurrent routers

Two goroutines calling `Route` at once race on `r.idx` — two messages can go to
the same worker and one can be skipped. Making it concurrent needs a mutex or
`atomic.AddUint64` on the cursor. The tests here are single-goroutine on purpose,
so the policy stays visible.
