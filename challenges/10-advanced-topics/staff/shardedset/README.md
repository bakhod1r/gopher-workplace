# A Set That Many Goroutines Can Share

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A deduplicating crawler guards its seen-set with one mutex. At thirty-two workers the mutex is the bottleneck, and a first attempt at striping reported the same URL as new twice.

## Task

Implement [shardedset.go](shardedset.go):

1. Insert `key` into its shard and report whether it was newly added.
2. The presence check and the insert must be one atomic step.
3. Hold only that shard's lock.
4. Correct under concurrent use: a key added by many goroutines reports true exactly once.

Replace the stub body in [shardedset.go](shardedset.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s.Add("a") twice
Output: true, then false
```

**Example 2:**

```
Input:  16 workers adding 200 shared keys
Output: exactly 200 trues
```

**Example 3:**

```
Input:  NewSet(0)
Output: still deduplicates
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Check-then-act must be atomic** | Releasing the lock between them lets two callers both win. |
| 2 | **Lock striping** | Independent shards mean independent locks. |
| 3 | **Deterministic routing** | The same key must always reach the same shard. |
| 4 | **Padding the shard** | Neighbouring mutexes on one line contend in hardware. |

## Hint

One lock, held across both the lookup and the insert.

## Validate

```bash
make verify
```
