# Drop Instead Of Blocking

**Level:** senior
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A request handler publishes a metric on a buffered channel. The collector stalls, the buffer fills, and every request in the service blocks behind a statistics counter.

## Task

Implement [dropwhenfull.go](dropwhenfull.go):

1. Send `v` on `ch` if it can be accepted immediately.
2. Report whether it was sent.
3. Never block, whatever the channel's state.

Replace the stub body in [dropwhenfull.go](dropwhenfull.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Offer(ch, 1) with room
Output: true
```

**Example 2:**

```
Input:  Offer on a full buffer
Output: false, the value is dropped
```

**Example 3:**

```
Input:  Offer on an unbuffered channel with no receiver
Output: false, immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **select with default** | Makes a channel operation non-blocking. |
| 2 | **Load shedding** | Dropping a sample beats stalling the request that produced it. |
| 3 | **Unbuffered means synchronous** | It can only accept when a receiver is already waiting. |

## Hint

One `select`, one case, one default.

## Validate

```bash
make verify
```
