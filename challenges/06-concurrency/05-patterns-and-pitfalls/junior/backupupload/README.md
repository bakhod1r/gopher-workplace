# Backup Upload Total

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The backup job uploads shards through a pool sized to the available bandwidth
and reports the total bytes transferred. The interesting choice here is how
the workers report back: instead of a shared counter behind a mutex, each
worker sends its number on a results channel and the caller adds them up.

## Task

Implement `TotalUploaded` in [backupupload.go](backupupload.go) so that:

1. It creates a jobs channel and a buffered results channel, and starts `workers` goroutines.
2. Each worker ranges over jobs and sends `upload(shard)` on the results channel.
3. After `close(jobs)`, `wg.Wait()` and `close(results)`, the caller sums the results and returns the total.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalUploaded([]string{"a", "bb"}, 2, size)
Output: 3
```

**Example 2:**

```
Input:  TotalUploaded([]string{"abcd"}, 1, size)
Output: 4
```

**Example 3:**

```
Input:  TotalUploaded(nil, 3, size)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Results channel** | Aggregating in one goroutine instead of sharing a counter. |
| 2 | **"Share memory by communicating"** | The channel carries ownership of each value to the summing goroutine. |
| 3 | **Close ordering** | close(jobs) → wg.Wait() → close(results), then drain and sum. |

## Hint

No mutex is needed: only the calling goroutine ever touches `total`, and the
workers only ever send.

## Validate

```bash
make verify
```
