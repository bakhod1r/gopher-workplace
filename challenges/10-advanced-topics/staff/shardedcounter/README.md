# Shard The Counter, Fold It Once

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A metrics counter behind one mutex is the bottleneck at sixteen cores. Replacing it with atomics per key is not possible — the keys are dynamic and the map itself needs protection.

## Task

Implement [shardedcounter.go](shardedcounter.go):

1. Increment `key`'s counter in the shard that owns it.
2. Hold only that shard's lock, and only for the update.
3. Correct under concurrent use: no lost increments, no race.

Replace the stub body in [shardedcounter.go](shardedcounter.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  c.Add("a", 1); c.Add("a", 2)
Output: Total()["a"] == 3
```

**Example 2:**

```
Input:  16 workers x 1000 adds over 8 keys
Output: 2000 per key
```

**Example 3:**

```
Input:  shardFor of one key
Output: always the same shard
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lock striping** | Splitting one lock into n reduces contention by roughly n. |
| 2 | **Deterministic sharding** | The same key must always map to the same shard, or increments split and are folded wrongly. |
| 3 | **Padding the shards** | Adjacent mutexes on one cache line reintroduce contention in hardware. |
| 4 | **Fold after the writers** | `Total` is a read-side operation, not part of the hot path. |

## Hint

One shard, one lock, one map update. The routing is already written.

## Validate

```bash
make verify
```
