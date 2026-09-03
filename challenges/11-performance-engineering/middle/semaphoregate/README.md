# A Permit To Proceed

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Not every limit is a worker pool. Sometimes the work arrives as goroutines you did not launch — inbound requests, callbacks — and what you need is a gate: at most N of you inside this section at once. A buffered channel is Go's semaphore, and the buffer size is the limit.

## Task

Implement the five pieces in [semaphoregate.go](semaphoregate.go):

1. `New(n)` builds a semaphore with `n` permits, defaulting a non-positive `n` to 1.
2. `Acquire` blocks for a permit, `TryAcquire` never blocks, and `Release` returns one.
3. An unmatched `Release` must be dropped rather than blocking or inflating the count; `Available` reports free permits.

## Examples

**Example 1:**

```
Input:  s := New(2); s.Acquire()
Output: Available 1
```

**Example 2:**

```
Input:  New(1); TryAcquire twice
Output: true, then false
```

**Example 3:**

```
Input:  New(2); Release three times
Output: Available still 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A buffered channel is a semaphore** | Sends block when it is full, which is exactly "wait for a permit". |
| 2 | **`select` with `default` never blocks** | That is the whole of `TryAcquire`. |
| 3 | **Bound the section, not the goroutines** | The callers already exist; the semaphore limits how many are inside at once. |

## Topics used again

Buffered channels, `select` with `default`, `len` and `cap` on channels.

## Hint

Decide whether a full buffer means "all permits taken" or "all permits free" — and keep `Available` consistent with it.

## Validate

```bash
make verify
```
