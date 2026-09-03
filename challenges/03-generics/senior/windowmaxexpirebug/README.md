# Sliding Maximum That Never Expires

**Level:** senior  
**Topic:** 03-generics

## Context

A rolling peak-latency chart keeps showing an old spike long after it has scrolled out of the window, so alerts never clear.

## Task

Fix the single planted bug in [windowmaxexpirebug.go](windowmaxexpirebug.go):

1. Find and fix the single bug so indices that fall out of the window are discarded.
2. The monotonic deque logic itself is correct.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  WindowMax([1,3,2], 2)
Output: [3 3]
```

**Example 2:**

```
Input:  WindowMax([5,1,1,1], 2)
Output: [5 1 1]
```

**Example 3:**

```
Input:  WindowMax([1], 5)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sliding window bookkeeping** | Every step must both admit the new index and evict the expired one. |
| 2 | **Monotonic deque** | The front holds the current maximum only while it is still inside the window. |

## Hint

What removes the front of the deque?

## Validate

```bash
make verify
```
