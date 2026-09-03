# A Stage That Stops When Told

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A pipeline stage is cancelled when the consumer hits an error. The cancellation unblocks the receives and the workers all sit blocked on their sends instead.

## Task

Implement [pipeline.go](pipeline.go):

1. Return a channel carrying each input doubled, computed by `workers` goroutines.
2. Close it once the input drains.
3. Every goroutine must exit when `done` is closed — whether it is receiving or sending.
4. `workers < 1` behaves as 1.

Replace the stub body in [pipeline.go](pipeline.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Stage(done, feed(1,2,3), 2)
Output: a channel yielding 2, 4 and 6 in some order
```

**Example 2:**

```
Input:  input drained
Output: out is closed
```

**Example 3:**

```
Input:  consumer abandons out, done closed
Output: no goroutine left behind
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Both ends can block** | Cancelling a receive is not enough if the send can block too. |
| 2 | **select on every channel operation** | Each one needs the `done` escape. |
| 3 | **WaitGroup then close** | The closer runs after all workers, in its own goroutine. |
| 4 | **Only one closer** | Workers must not close a channel they share. |

## Hint

Two blocking operations per iteration. Both need a way out.

## Validate

```bash
make verify
```
