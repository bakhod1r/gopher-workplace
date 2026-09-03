# The Minimum That Goes Stale

**Level:** senior  
**Topic:** 03-generics

## Context

A constant-time minimum stack keeps reporting a minimum that was popped several operations ago.

## Task

Fix the single planted bug in [minstacksyncbug.go](minstacksyncbug.go):

1. Find and fix the single bug so the two stacks stay in step.
2. `Min` must reflect only the elements still present.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Push(3); Push(1); Pop(); Min()
Output: 3
```

**Example 2:**

```
Input:  Push(3); Min()
Output: 3
```

**Example 3:**

```
Input:  Pop() on empty
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Structural invariants** | Every operation must restore what the type promises about itself. |
| 2 | **Paired state** | Two stacks that encode one thing must be updated together. |

## Hint

`Push` maintains `mins`. What maintains it on the way out?

## Validate

```bash
make verify
```
