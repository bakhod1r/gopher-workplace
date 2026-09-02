# Two Stacks, One Wrong Transfer

**Level:** senior  
**Topic:** 03-generics

## Context

A queue built from two stacks serves elements out of order as soon as producers and consumers interleave.

## Task

Fix the single planted bug in [queuetwostacksbug.go](queuetwostacksbug.go):

1. Find and fix the single bug so FIFO order survives interleaved enqueues and dequeues.
2. An empty queue must still report `false`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Enq 1,2; Deq; Enq 3; Deq; Deq
Output: 1, 2, 3
```

**Example 2:**

```
Input:  Enq 1; Deq
Output: 1
```

**Example 3:**

```
Input:  Deq on empty
Output: zero, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Order of operations** | Doing the right steps in the wrong order is still a bug. |
| 2 | **Transfer only when drained** | Pouring new items onto a partly drained out-stack buries the older ones. |

## Hint

Under what condition may the in-stack be transferred?

## Validate

```bash
make verify
```
