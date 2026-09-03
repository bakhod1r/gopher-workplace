# Splitting The Lock

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A single mutex over a whole map is a queue with a nice name: every goroutine waits for every other, however unrelated their keys. Striping cuts the map into shards with independent locks, so contention drops by roughly the shard count — the design behind `sync.Map`'s read path, Java's old `ConcurrentHashMap`, and most sharded caches.

## Task

Implement the six pieces in [striped.go](striped.go):

1. `New(n)` builds `n` shards (a non-positive `n` gives 1), and `Shards` reports the count.
2. `ShardOf` maps a key to a shard index, stably for the lifetime of the map and spread across the shards.
3. `Set` and `Get` take only their own shard's lock; `Len` totals every shard, and the whole thing must be race-free.

## Examples

**Example 1:**

```
Input:  m := New(8); Set("a", 1); Get("a")
Output: 1, true
```

**Example 2:**

```
Input:  ShardOf("a") called 100 times
Output: the same index every time
```

**Example 3:**

```
Input:  1000 distinct keys over 8 shards
Output: at least 6 shards used
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Contention scales with sharing** | Fewer goroutines per lock means less waiting, almost linearly. |
| 2 | **The hash must be stable** | A key that moves between shards is a key that gets lost. |
| 3 | **Cross-shard operations are expensive** | `Len` touches every lock, which is why sharded maps rarely offer one. |

## Topics used again

`sync.RWMutex`, hashing with `hash/maphash`, slices of structs, modular arithmetic.

## Hint

`maphash.String(m.seed, key)` gives a `uint64`; the shard is that modulo the shard count. Index the shard slice rather than ranging it by value — a `sync.RWMutex` must not be copied.

## Validate

```bash
make verify
```
