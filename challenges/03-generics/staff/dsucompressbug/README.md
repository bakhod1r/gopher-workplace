# The Union-Find That Walks The Chain Forever

**Level:** staff  
**Topic:** 03-generics

## Context

A connectivity service answers every query correctly, so it shipped. Then a customer imported a dataset whose merges form one long chain, and p99 latency went from microseconds to seconds. Nothing in the answers is wrong — only the cost.

## Task

Fix the single planted bug in [dsucompressbug.go](dsucompressbug.go):

1. Find and fix the single bug so `Find` flattens the path it walks.
2. The answers must not change: the same elements must stay connected.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Union(1,2); Union(2,3); Connected(1,3)
Output: true
```

**Example 2:**

```
Input:  Connected(1,9) with 9 unseen
Output: false
```

**Example 3:**

```
Input:  15000 chained unions, then 15000 queries
Output: completes well under the time budget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Failures that need scale** | A defect that small inputs cannot express is still a defect; test at size. |
| 2 | **Amortised cost** | Path compression turns a linear walk into a near-constant one by rewriting what it traverses. |
| 3 | **Correct but slow is still broken** | At staff level the complexity class is part of the contract. |

## Hint

The buggy loop finds the right root. What does it fail to leave behind?

## Validate

```bash
make verify
```
