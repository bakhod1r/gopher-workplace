# Breadth-First Order

**Level:** middle  
**Topic:** 03-generics

## Context

A dependency viewer lists what a task pulls in, nearest first, and the graph may contain cycles.

## Task

Implement the stub(s) in [bfsgen.go](bfsgen.go):

1. Implement `BFS`, returning reachable nodes in breadth-first order starting with `start`.
2. Visit each node once, even when the graph has cycles.
3. Neighbours are visited in their stored order.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BFS({a:[b,c], b:[d]}, a)
Output: [a b c d]
```

**Example 2:**

```
Input:  BFS({a:[a]}, a)
Output: [a]
```

**Example 3:**

```
Input:  BFS({}, a)
Output: [a]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Queue-driven traversal** | A FIFO queue is what makes the order breadth-first rather than depth-first. |
| 2 | **Marking on enqueue** | Marking when queued — not when visited — prevents duplicates in cyclic graphs. |
| 3 | **Deterministic output** | Order follows the adjacency lists, so the result is reproducible. |

## Hint

Mark nodes as seen when you enqueue them, not when you dequeue them.

## Validate

```bash
make verify
```
