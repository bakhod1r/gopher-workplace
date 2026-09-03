# Reservation Thrown Away

**Level:** senior  
**Topic:** 03-generics

## Context

A batch writer silently drops every record it was asked to add, while the reservation code around it looks correct.

## Task

Fix the single planted bug in [growbug.go](growbug.go):

1. Find and fix the single bug so the appended values reach the caller.
2. The result must contain the original elements followed by the new ones.
3. Keep the reservation — it is what makes the appends allocate once.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Collect([]int{1}, 2, 3)
Output: []int{1,2,3}
```

**Example 2:**

```
Input:  cap after the call
Output: >= 3
```

**Example 3:**

```
Input:  Collect(nil, 1)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value semantics** | `Grow` returns a new slice header; it cannot change the caller's variable. |
| 2 | **Ignored return values** | The compiler does not complain, so the reservation is silently discarded. |
| 3 | **Same shape as `append`** | Both return a slice you must keep. |

## Hint

Which slice header does the function return, and which one was appended to?

## Validate

```bash
make verify
```
