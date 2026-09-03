# Block The Producer, Do Not Buffer Forever

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A producer sends into an unbounded queue whenever the consumer lags. The lag is normal and the queue is the reason the process runs out of memory during the nightly batch.

## Task

Implement [boundedqueue.go](boundedqueue.go):

1. Append `v`, waiting while the queue is full.
2. Report false without appending when `done` closes first.
3. The queue must never hold more than its capacity.

Replace the stub body in [boundedqueue.go](boundedqueue.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewQueue(2), two Puts
Output: true, true
```

**Example 2:**

```
Input:  a third Put while full
Output: blocks, then false when done closes
```

**Example 3:**

```
Input:  a Take while a Put waits
Output: the Put proceeds
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Backpressure by blocking** | A full buffer must slow the producer, not grow. |
| 2 | **Buffered channel as the queue** | Its capacity is the bound and its blocking is the pressure. |
| 3 | **Cancellable waiting** | `select` over the send and `done`. |
| 4 | **Append nothing on cancellation** | Returning false must leave the queue untouched. |

## Hint

A `select` with the send and the cancellation.

## Validate

```bash
make verify
```
