# Bounded Blocking Queue

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An ingest queue must apply backpressure: producers block when it is full, consumers block when it is empty, and a close must wake everyone.

## Task

Implement the stub(s) in [boundedq.go](boundedq.go):

1. Implement `Push`, `Pop`, `Close`, and `Len` on `*Queue` using a mutex and condition variables.
2. `Push` after `Close` returns false; `Pop` drains remaining items, then returns false once empty and closed.
3. Constraint: `-race` clean, no busy-waiting, no deadlock, and `Close` must wake every blocked goroutine.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  capacity 1; Push(1), then Pop()
Output: 1, true
```

**Example 2:**

```
Input:  Close with items still queued
Output: Pop drains them, then reports false
```

**Example 3:**

```
Input:  Push after Close
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Condition variables** | `sync.Cond` blocks without spinning and wakes on state change. |
| 2 | **Broadcast on close** | Every waiter must re-check the predicate and exit. |
| 3 | **Predicate loops** | `for !cond { wait() }` — never `if`, because wakeups are not promises. |

## Hint

Two conds sharing one mutex: `notFull` for producers, `notEmpty` for consumers. `Close` broadcasts both.

## Validate

```bash
make verify
```
