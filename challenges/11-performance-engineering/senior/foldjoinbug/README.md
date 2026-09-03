# A Flame Graph With Impossible Call Paths

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The flame graph shows `encode` calling `handler`, which never happens — the call goes the other way. Two genuinely different paths have also been fused into one very wide bar. Something is treating the stack as a set of functions rather than as an ordered path.

## Task

Fix the single planted bug in [foldjoinbug.go](foldjoinbug.go):

1. Find and fix the one bug so each line preserves the stack's original frame order.
2. Stacks with the same frames in a different order must remain distinct.
3. Identical stacks must still be summed, and the existing ordering of the output must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Fold([{[main a] 3}])
Output: ["main;a 3"]
```

**Example 2:**

```
Input:  Fold([{[a b] 1} {[b a] 2}])
Output: ["b;a 2" "a;b 1"]
```

**Example 3:**

```
Input:  Fold([{[main rec rec] 4}])
Output: ["main;rec;rec 4"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A stack is a sequence** | The order *is* the caller-callee relationship; sorting it destroys the data. |
| 2 | **Aggregating by the wrong key** | Sorted frames merge distinct paths into one plausible-looking line. |
| 3 | **Impossible edges are a tell** | A graph showing a callee calling its caller is a key-construction bug. |

## Hint

The frames are being normalised before they become the map key.

## Validate

```bash
make verify
```
