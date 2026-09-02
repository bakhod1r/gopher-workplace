# Fan-Out Sink

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

Every event must reach several sinks. One slow or broken sink must not stop the others, and the fan-out runs concurrently.

## Task

Implement the stub(s) in [sinkfanout.go](sinkfanout.go):

1. Implement `Write` on `*MemSink` (append) and on `ErrSink` (always fail).
2. Implement `FanOut`, which writes an event to every sink concurrently and returns the number of failures.
3. Constraint: race-free under `-race`, and every sink must be attempted even when some fail.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FanOut with 2 good sinks and 1 failing
Output: 1 failure, both good sinks received the event
```

**Example 2:**

```
Input:  FanOut with no sinks
Output: 0
```

**Example 3:**

```
Input:  100 sinks
Output: all 100 receive the event
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Concurrent fan-out** | Independent sinks are naturally parallel. |
| 2 | **Race-free counters** | `sync/atomic` or per-goroutine results, never a bare `n++`. |
| 3 | **Isolation of failures** | Reused: one bad implementer must not abort the loop. |

## Hint

Count failures with `atomic.AddInt64` or collect into a preallocated slice indexed per sink.

## Validate

```bash
make verify
```
