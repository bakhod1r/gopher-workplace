# The Cache That Walks Itself

**Level:** staff  
**Topic:** 03-generics

## Context

A hot-path LRU cache in front of a slow store is supposed to be free. Under load its own CPU time dominates the profile, and the cost per lookup rises as the cache warms up.

## Task

Fix the single planted bug in [lruo1bug.go](lruo1bug.go):

1. Find and fix the single bug so promotion is constant time.
2. The recency order and the return values must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Put(1,"a"); Get(1)
Output: "a", true
```

**Example 2:**

```
Input:  cap 2: Put(1),Put(2),Get(1),Put(3); Keys()
Output: [3 1]
```

**Example 3:**

```
Input:  150000 Gets on a 16384-entry cache
Output: well under the time budget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Asymptotic cost** | A per-operation cost that grows with the container size is a production outage waiting to happen. |
| 2 | **Linked lists exist for O(1) splicing** | Unlink and push-front touch four pointers; rebuilding touches every node. |
| 3 | **Amortisation does not save you** | The rebuild is paid in full on every single lookup, not occasionally. |

## Hint

How much work does one `Get` do when the cache holds `n` entries?

## Validate

```bash
make verify
```
