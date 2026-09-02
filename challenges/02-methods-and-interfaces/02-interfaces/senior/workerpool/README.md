# Worker Pool

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A batch job runs tasks through a fixed pool so a burst of work cannot spawn a million goroutines.

## Task

Implement the stub(s) in [workerpool.go](workerpool.go):

1. Implement `Run` on `SquareTask`.
2. Implement `RunAll`, which executes tasks with exactly `workers` goroutines and returns results in input order.
3. Constraint: the pool must be race-free under `-race`, and must never start more than `workers` goroutines.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RunAll(tasks, 4) with 100 tasks
Output: 100 results, input order preserved
```

**Example 2:**

```
Input:  RunAll(tasks, 1)
Output: same results, sequential
```

**Example 3:**

```
Input:  RunAll(nil, 4)
Output: empty
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded concurrency** | Worker count is the resource ceiling, not the task count. |
| 2 | **Index-addressed results** | Writing to distinct slice slots needs no mutex. |
| 3 | **sync.WaitGroup** | Reused from concurrency: wait for every worker to finish. |

## Hint

Each worker takes indexes off one channel and writes `out[i]` — distinct indexes, no lock needed.

## Validate

```bash
make verify
```
