# First Replica to Answer Wins

**Level:** middle
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

The read path serves the same row from three replicas. Latency is dominated by the slowest one, so the query is issued to all of them and the first answer wins. What makes this safe rather than wasteful is the shared derived context: the winner's return cancels it, and the losing replicas abandon their scans instead of finishing work that will be thrown away.

## Task

Implement the exported function(s) in [replicafanout.go](replicafanout.go) so that:

1. It returns `ErrNoReplicas` for an empty replica set, and `ctx.Err()` — without calling any replica — when the request context is already finished.
2. It derives one cancellable context from `ctx`, queries every replica concurrently with it, and `defer`s the cancel.
3. It returns the first successful row; the remaining replicas must observe `context.Canceled`.
4. If every replica fails it returns `"", failures[0]` — the error from the replica at index 0, regardless of arrival order.
5. It must be race-free under `-race` and must not leak goroutines that block forever on an unbuffered channel.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ReadFromReplicas(ctx, nil)
Output: "", ErrNoReplicas
```

**Example 2:**

```
Input:  ReadFromReplicas(ctx, [blocked, fast("row-2"), blocked])
Output: "row-2", nil  (both blocked replicas see context.Canceled)
```

**Example 3:**

```
Input:  ReadFromReplicas(cancelled ctx, [fast])
Output: "", context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-out with a shared child context** | One `WithCancel` covers every replica; cancelling it stops all of them. |
| 2 | **Buffered result channel** | Sized `len(replicas)` so losing goroutines can always send and exit, never leak. |
| 3 | **Deterministic error selection** | Index the failures by replica position instead of trusting arrival order. |

## Hint

Give each goroutine a buffered slot to report into, and record failures at `answers[i]` so "which error do we return" does not depend on scheduling.

## Validate

```bash
make verify
```
