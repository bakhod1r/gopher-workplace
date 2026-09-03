# Every Profile Says `main` Is Hot

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

The flat column shows `main` at 100% and nothing else above a rounding error. That is exactly what a cumulative column looks like — except this is supposed to be self time, and self time is the one column `main` should be absent from.

## Task

Fix the single planted bug in [selftimeleakbug.go](selftimeleakbug.go):

1. Find and fix the one bug so only the frame that was executing is credited.
2. The same function must accumulate self time across every stack it terminates.
3. The existing junk-sample guards must keep working.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SelfTime([{[main a b] 5}])
Output: {b:5}
```

**Example 2:**

```
Input:  SelfTime([{[main a] 3} {[main b a] 4} {[main b] 1}])
Output: {a:7 b:1}
```

**Example 3:**

```
Input:  SelfTime([{[main handler parse] 10} {[main handler encode] 20}])
Output: {parse:10 encode:20}, no main
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Callers are waiting, not working** | At the instant of the sample only the leaf holds the CPU. |
| 2 | **Stack order is a convention** | Caller-first or leaf-first — pick wrong and every attribution inverts. |
| 3 | **`main` at the top of a flat profile** | Almost always means the leaf and the root got swapped. |

## Hint

Which end of the stack is the function that was running?

## Validate

```bash
make verify
```
