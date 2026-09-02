# Breadth-First With Duplicates

**Level:** senior  
**Topic:** 03-generics

## Context

A dependency walker over a diamond-shaped graph reports the join node twice, and downstream steps run twice with it.

## Task

Fix the single planted bug in [bfsmarkbug.go](bfsmarkbug.go):

1. Find and fix the single bug so every reachable node is visited once.
2. The order must stay breadth-first.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  diamond a->b,c; b,c->d
Output: [a b c d]
```

**Example 2:**

```
Input:  d appears
Output: once
```

**Example 3:**

```
Input:  isolated node
Output: [n]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 2 | **Mark on enqueue** | Marking on dequeue lets a node be queued many times before it is first popped. |

## Hint

When is a node recorded as seen?

## Validate

```bash
make verify
```
