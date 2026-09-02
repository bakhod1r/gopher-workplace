# Deadline Guard

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A batch step must finish before a deadline. Rather than being killed mid-flight, it stops cleanly and reports what it completed.

## Task

Implement the stub(s) in [deadlineifc.go](deadlineifc.go):

1. Implement `Do` on `*CountingOp`.
2. Implement `RunUntil`, which runs the operation repeatedly until the deadline passes, returning the completed count.
3. Constraint: check the deadline before each call — never start work that cannot be attributed, and never sleep past the deadline.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  a deadline already passed
Output: 0 operations run
```

**Example 2:**

```
Input:  a clock advancing one unit per operation, 5 units of budget
Output: 5 operations
```

**Example 3:**

```
Input:  a zero-cost operation with a fixed budget
Output: bounded by the budget
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deadline propagation** | Time budget is an input, not an ambient global. |
| 2 | **Clock injection** | Reused: deterministic tests without sleeping. |
| 3 | **Pre-check discipline** | Reused: stop before starting work you cannot finish. |

## Hint

`for !clock.Now().Before(deadline) { break }` — or equivalently, loop while `Now().Before(deadline)`.

## Validate

```bash
make verify
```
