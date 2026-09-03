# Three Counters, Three Costs

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A mutex-guarded counter serialises every writer through one lock. An atomic counter drops the lock but keeps the shared word, so the cache line still ping-pongs between cores. A sharded counter removes the sharing itself — and pays for it with an approximate read. Same answer, three very different scaling curves.

## Task

Implement all three counters in [mutexvsatomic.go](mutexvsatomic.go):

1. `MutexCounter` and `AtomicCounter` both count correctly under concurrency.
2. `NewSharded(n)` builds `n` shards (a non-positive `n` gives 1), and `Inc(id)` picks a shard from `id` without panicking on a negative one.
3. `Value` on the sharded counter sums the shards; all three must agree after the writers finish.

## Examples

**Example 1:**

```
Input:  MutexCounter: Inc(), Inc()
Output: 2
```

**Example 2:**

```
Input:  NewSharded(4): Inc(0), Inc(1), Inc(9)
Output: 3
```

**Example 3:**

```
Input:  50 goroutines × 2000 increments on each counter
Output: all three report 100000
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **The lock is not the only cost** | Atomics remove the lock but not the contended cache line. |
| 2 | **Sharding removes the sharing** | Different writers, different words — the only approach that scales with cores. |
| 3 | **A sharded read is approximate** | The shards are read at different instants, so the total is only exact when writers stop. |

## Topics used again

`sync.Mutex`, `sync/atomic`, slices of atomics, modular arithmetic.

## Hint

A negative `id` needs handling before the modulo: `((id % n) + n) % n`, or take the absolute value first.

## Validate

```bash
make verify
```
