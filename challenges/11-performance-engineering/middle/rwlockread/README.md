# Many Readers, One Writer

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

Configuration, routing tables, feature flags: read on every request, written a handful of times a day. A plain mutex makes every one of those reads wait for the others, for no reason at all — readers do not conflict. `sync.RWMutex` lets them run together and only excludes them while a writer holds the lock.

## Task

Implement the four methods in [rwlockread.go](rwlockread.go):

1. `Get` and `Version` read under a read lock; the zero `Config` works with no constructor.
2. `Replace` swaps in a copy of the given map under the write lock and bumps the version.
3. `Snapshot` returns the values and their version together, as a consistent pair, with an empty non-nil map when nothing is set.

## Examples

**Example 1:**

```
Input:  Replace({k:v}); Get("k")
Output: "v", true; Version 1
```

**Example 2:**

```
Input:  in := {k:v}; Replace(in); in["k"] = "mutated"
Output: Get("k") is still "v"
```

**Example 3:**

```
Input:  Snapshot(), then Replace({b:2})
Output: the snapshot still holds only a
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`RLock` does not exclude readers** | Only writers are excluded, so a read-mostly workload stops serialising. |
| 2 | **Copy on write, copy on snapshot** | The live map must never be reachable from outside the lock. |
| 3 | **Consistency is per lock acquisition** | Reading the map and the version separately can straddle a write. |

## Topics used again

`sync.RWMutex`, defensive map copies, `defer`, multiple return values.

## Hint

Calling `c.Version()` from inside `Snapshot` would take the read lock twice — take it once and read both fields.

## Validate

```bash
make verify
```
