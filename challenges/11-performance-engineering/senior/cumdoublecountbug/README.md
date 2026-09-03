# The Recursive Function That Owns 4000% Of The CPU

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

A profile of a recursive descent parser shows one function with a cumulative time forty times the program's wall clock. Everyone assumes the profiler is broken. The seen-set is right there in the loop; it is just not being consulted.

## Task

Fix the single planted bug in [cumdoublecountbug.go](cumdoublecountbug.go):

1. Find and fix the one bug so a function appearing several times in one stack is credited once for that sample.
2. Frames must still be credited once per *sample* they appear in, across the whole profile.
3. The existing junk-sample guards must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  CumSum([{[main rec rec rec] 6}])
Output: {main:6 rec:6}
```

**Example 2:**

```
Input:  CumSum([{[a b a b] 4}])
Output: {a:4 b:4}
```

**Example 3:**

```
Input:  CumSum([{[main a] 3} {[main b] 4}])
Output: {main:7 a:3 b:4}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cum is per sample, not per frame** | A stack is a set of responsible functions, however deep the recursion goes. |
| 2 | **Recording without checking** | Marking the set and never reading it is a very quiet no-op. |
| 3 | **Absurd percentages are aggregation bugs** | A cum above the wall clock is possible; forty times it is not. |

## Hint

The seen-set is being written to. Nothing reads it.

## Validate

```bash
make verify
```
