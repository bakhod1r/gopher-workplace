# Parallel Level Counts

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The nightly log report is produced by splitting a day's log file into chunks
and counting severity levels in each one concurrently. Go maps are not safe
for concurrent writes — the runtime will kill the process with a "concurrent
map writes" fatal error — so the merge has to be guarded.

## Task

Implement `CountLevels` in [logcounts.go](logcounts.go) so that:

1. It starts one goroutine per chunk, tracked by a `WaitGroup`.
2. Each goroutine tallies its chunk into a *local* map, with no locking.
3. It then takes the mutex once and folds the local counts into the shared totals; the totals are returned after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CountLevels([][]string{{"ERR a"}, {"ERR b"}}, level)
Output: map[ERR:2]
```

**Example 2:**

```
Input:  CountLevels([][]string{{"ERR a", "INFO b"}}, level)
Output: map[ERR:1 INFO:1]
```

**Example 3:**

```
Input:  CountLevels(nil, level)
Output: empty map
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Mutex-guarded map** | Concurrent map writes are a fatal runtime error, not just a race. |
| 2 | **Local-then-merge** | Do the bulk work lock-free, then hold the lock for one short merge. |
| 3 | **WaitGroup** | The totals may only be read after every chunk has merged. |

## Hint

Count into a map the goroutine owns, then lock once and add the local counts
into the shared map — one lock acquisition per chunk, not per line.

## Validate

```bash
make verify
```
