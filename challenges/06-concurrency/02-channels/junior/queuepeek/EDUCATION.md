# Non-Blocking Receive With default

## Intuition

Adding `default` to a `select` removes all blocking: if no case can proceed
this instant, `default` runs. "Can proceed" includes receiving from a
*closed* channel, which is why a peek at a closed queue reports
`0, true` — the receive really did happen, it just produced the zero value.

## Approach

1. Write a `select` with one receive case and a `default`.
2. In the receive case return the job id and `true`.
3. In `default` return `0` and `false`.

## Solution

```go
func PeekJob(jobs <-chan int) (int, bool) {
	select {
	case id := <-jobs:
		return id, true
	default:
		return 0, false
	}
}
```

## Walkthrough

With job `5` queued, the receive case is ready and wins: `5, true`. On an
empty open queue nothing is ready, so `default` fires: `0, false`. On a
closed queue the receive is ready and yields the zero value: `0, true`.

## Pitfalls

- Expecting `false` for a closed queue — use the comma-ok form inside the case if you need to tell those apart.
- Polling in a tight loop burns CPU; blocking receives are usually right.
- `default` in a `select` is not the same as a `case` — it never waits.
