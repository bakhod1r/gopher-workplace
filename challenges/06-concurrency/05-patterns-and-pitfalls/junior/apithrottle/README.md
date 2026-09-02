# Throttled API Calls

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The upstream API returns 429 above a fixed concurrency, so the client throttles
itself with a semaphore. This version also instruments itself: it tracks how
many calls were genuinely in flight at once, which is the number an operator
wants when tuning the limit.

## Task

Implement `PeakInFlight` in [apithrottle.go](apithrottle.go) so that:

1. It gates every request goroutine on a buffered channel of capacity `limit`.
2. Under a mutex it increments an in-flight counter on entry, records a new peak when one is reached, and decrements on exit.
3. It returns the peak after `wg.Wait()`; the peak is 0 for no requests and never exceeds `limit`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PeakInFlight(5 requests, limit 2, do)
Output: 1 or 2, never more
```

**Example 2:**

```
Input:  PeakInFlight(3 requests, limit 1, do)
Output: 1
```

**Example 3:**

```
Input:  PeakInFlight(nil, limit 4, do)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Semaphore limit** | Capacity is a hard ceiling: the peak can never exceed it. |
| 2 | **Instrumenting concurrency** | In-flight counters must be updated under the same lock as the peak. |
| 3 | **Non-deterministic timing** | The exact peak depends on scheduling; only the bound is guaranteed. |

## Hint

Increment and compare against the peak inside one critical section — reading
the counter and updating the peak separately can record a value that never
actually happened.

## Validate

```bash
make verify
```
