# Scrape Queue Stats

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The Prometheus-style scrape agent sizes its sample queue at start-up. Before
going live it fills the queue to capacity and reports the depth and the
configured size so the operator can confirm the sizing.

## Task

Implement `QueueStats` in [scrapequeue.go](scrapequeue.go) so that:

1. It creates a buffered channel with capacity `size`.
2. It sends exactly `size` samples without blocking.
3. It returns `len(queue)` then `cap(queue)`; `size <= 0` gives `0, 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  QueueStats(3)
Output: 3, 3
```

**Example 2:**

```
Input:  QueueStats(0)
Output: 0, 0
```

**Example 3:**

```
Input:  QueueStats(-1)
Output: 0, 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Buffered channel** | `make(chan T, n)` queues up to `n` samples. |
| 2 | **`len` vs `cap`** | `len` = samples waiting now; `cap` = configured queue size. |
| 3 | **Blocking rule** | The `size+1`-th send blocks until a scraper receives. |

## Hint

`len` and `cap` work on channels exactly as they do on slices. Fill the
queue with a plain `for` loop.

## Validate

```bash
make verify
```
