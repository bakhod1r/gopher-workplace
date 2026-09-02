# Error Group

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A fan-out of sub-requests must stop as soon as one fails, report the first error, and leave no goroutine behind.

## Task

Implement the stub(s) in [errgrouplite.go](errgrouplite.go):

1. Implement `Go` and `Wait` on `*Group`, running tasks with at most `Limit` concurrent goroutines.
2. The first error is remembered and cancels the group; later errors are discarded.
3. Constraint: `-race` clean, `Wait` returns only after every started goroutine has finished, and the goroutine count returns to baseline.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  three tasks, the second fails
Output: Wait returns that error
```

**Example 2:**

```
Input:  all tasks succeed
Output: nil
```

**Example 3:**

```
Input:  Limit 2 with 100 tasks
Output: peak concurrency 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **errgroup semantics** | First error wins; cancellation is the group's job, not the caller's. |
| 2 | **Bounded fan-out** | Reused: a semaphore acquired before spawning bounds goroutines too. |
| 3 | **Cancellation signal** | A closed channel tells running tasks to stop early. |

## Hint

`sync.Once` to record the first error and close the cancel channel exactly once.

## Validate

```bash
make verify
```
