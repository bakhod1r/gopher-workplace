# The Self-Constraint Called Backwards

**Level:** staff  
**Topic:** 03-generics

## Context

A release picker that should roll back to the oldest healthy version keeps selecting the newest one. The comparison type has its own `Less` method and its unit tests all pass.

## Task

Fix the single planted bug in [lesserselfbug.go](lesserselfbug.go):

1. Find and fix the single bug so the smallest element is returned.
2. An empty input must still report `false`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  MinOf([]ver{{1,5},{1,2},{2,0}})
Output: ver{1,2}, true
```

**Example 2:**

```
Input:  MinOf([]ver{{0,1},{9,9}})
Output: ver{0,1}, true
```

**Example 3:**

```
Input:  MinOf([]ver{})
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Self-referential constraints** | `Lesser[T]` names a method taking the very type being constrained, so both sides of the call have the same type. |
| 2 | **Receiver and argument are not symmetric** | `a.Less(b)` and `b.Less(a)` are opposite questions that compile identically. |
| 3 | **Silent direction errors** | A reversed comparison is never a type error — only a wrong answer. |

## Hint

Which of the two values should be the receiver?

## Validate

```bash
make verify
```
