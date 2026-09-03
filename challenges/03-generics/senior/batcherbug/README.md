# The Batch That Never Ships

**Level:** senior  
**Topic:** 03-generics

## Context

An upload pipeline batches records by 100. Every run silently loses up to 99 records at the tail.

## Task

Fix the single planted bug in [batcherbug.go](batcherbug.go):

1. Find and fix the single bug so a partial final batch is still emitted.
2. Full batches must keep their exact size.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Batches([]int{1,2,3}, 2)
Output: [][]int{{1,2},{3}}
```

**Example 2:**

```
Input:  Batches([]int{1,2,3,4}, 2)
Output: [][]int{{1,2},{3,4}}
```

**Example 3:**

```
Input:  Batches([]int{1}, 0)
Output: [][]int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Flushing the accumulator** | A loop that only emits on a full buffer must flush after the loop. |
| 2 | **Off-by-a-batch** | The failure is invisible whenever the input divides evenly by the batch size. |

## Hint

What happens to `cur` when the loop ends?

## Validate

```bash
make verify
```
