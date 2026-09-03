# Read, Modify, Write — Three Chances To Lose

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

`m[k]++` looks atomic and is not: it reads, adds, and writes back. Fifty goroutines doing that to the same key lose tens of thousands of increments, and a Go map does not even survive concurrent writes — it detects them and kills the process. The fix is a mutex around the whole read-modify-write, not around each half.

## Task

Implement the three methods in [sharedcounter.go](sharedcounter.go):

1. `Add` increases a named counter by `delta`, losing no updates under concurrency.
2. `Get` returns one counter's value, `0` for an unknown key.
3. `Snapshot` returns a copy the caller can use without the lock; the zero `Counter` must work without a constructor.

## Examples

**Example 1:**

```
Input:  Add("a", 1); Add("a", 2)
Output: Get("a") is 3
```

**Example 2:**

```
Input:  50 goroutines each adding 1000
Output: exactly 50000
```

**Example 3:**

```
Input:  s := Snapshot(); s["a"] = 999
Output: the counter is unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`++` is three operations** | Read, add, store — an interleaving between any two loses an update. |
| 2 | **The lock covers the whole operation** | Locking the read and the write separately is not the same as locking the update. |
| 3 | **Snapshots need the lock too** | Ranging a map while another goroutine writes it is a race and, for maps, a crash. |

## Topics used again

`sync.Mutex`, lazy map initialisation, defensive copies, `defer`.

## Hint

`defer c.mu.Unlock()` immediately after `Lock`, then treat the rest of the method as a critical section.

## Validate

```bash
make verify
```
