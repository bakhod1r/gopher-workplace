# Panic And Unwind

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A worker panicked and took the process down. Recovery has to happen on the panicking goroutine, in the right defer, without swallowing the diagnosis.

## Task

Implement the stub(s) in [panicunwind.go](panicunwind.go):

1. Implement `SafeRun`, which runs a task and converts a panic into an error.
2. Implement `RunAll`, which runs every task and returns the recovered errors in order, with nil for the tasks that succeeded.
3. Implement `Order`, showing that deferred calls run last-in-first-out even while unwinding.
4. Constraint: no panic escapes `SafeRun`, a nil panic value is still reported, and a `runtime.Error` keeps its message.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SafeRun on a task that panics with "boom"
Output: an error mentioning boom
```

**Example 2:**

```
Input:  a task that returns an error normally
Output: that error, unchanged
```

**Example 3:**

```
Input:  Order()
Output: the deferred calls in LIFO order
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **recover semantics** | `recover` only works when called directly from a deferred function of the panicking goroutine. |
| 2 | **Named results** | A deferred closure can rewrite a named result after a panic. |
| 3 | **Defer ordering** | Deferred calls run LIFO, including during unwinding. |

## Hint

Use a named `err` result so the deferred `recover` can assign to it.

## Validate

```bash
make verify
```
