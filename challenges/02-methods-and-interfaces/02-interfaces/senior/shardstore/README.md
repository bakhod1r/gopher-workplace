# Sharded Store

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A single mutex around one map became the bottleneck under concurrent load. The store is now sharded by key.

## Task

Implement the stub(s) in [shardstore.go](shardstore.go):

1. Implement `shardFor`, `Put`, `Get`, and `Len` on `*ShardedStore`.
2. Each shard has its own mutex so unrelated keys do not contend.
3. Constraint: race-free under `-race`, correct under concurrent writers, and a key must always map to the same shard.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Put("a", "1"); Get("a")
Output: "1", true
```

**Example 2:**

```
Input:  1000 concurrent writers
Output: every key readable, Len == 1000
```

**Example 3:**

```
Input:  the same key twice
Output: the same shard
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lock sharding** | Contention drops roughly linearly with shard count. |
| 2 | **Deterministic hashing** | The same key must always land in the same shard. |
| 3 | **sync.Mutex per shard** | Reused: fine-grained locking instead of one global lock. |

## Hint

A simple FNV-style byte hash modulo the shard count is enough — it just has to be deterministic.

## Validate

```bash
make verify
```
