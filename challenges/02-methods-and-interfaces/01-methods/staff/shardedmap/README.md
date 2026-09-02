# Sharded Map

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

One mutex over one map serializes every writer, however unrelated their keys. A
sharded map splits the data across N independently locked maps, so two writers
collide only when their keys hash to the same shard — contention drops roughly
N-fold. `getShard` already does the hashing.

## Task

Implement `Set` and `Get` on `*ShardedMap` in [shardedmap.go](shardedmap.go):

1. `Set` selects the shard, takes its **write** lock, stores, releases.
2. `Get` selects the shard, takes its **read** lock, returns the value and the
   comma-ok flag.

**Constraint (staff):** 8 goroutines × 5,000 mixed read/write operations must stay correct and `-race`-clean, and keys must spread across every shard.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Set("a", 1); Get("a")
Output: (1, true)
```

**Example 2:**

```
Input:  Get("missing")
Output: (0, false)
```

**Example 3:**

```
Input:  1000 goroutines each Set a distinct key
Output: Len() == 1000, every key readable, clean under -race
```

_Explanation:_ distinct keys mostly land on distinct shards and proceed in parallel.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Lock granularity** | The lock belongs to the shard, not to the map — that is the entire optimization. |
| 2 | **Read vs write locks** | `Get` must use `RLock` or the sharding buys much less. |
| 3 | **Hash-based routing** | The same key must always reach the same shard, or reads miss writes. |

## Hint

Two lines of body each, plus the lock pair. Lock the shard you selected — never
a map-wide lock, and never a different shard than the one you write to.

## Validate

```bash
make verify
```
