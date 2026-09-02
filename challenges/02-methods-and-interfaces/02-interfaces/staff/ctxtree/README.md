# Cancellation Tree

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A request spawns sub-operations. Cancelling the request must cancel every descendant, and cancelling a child must not touch its siblings.

## Task

Implement the stub(s) in [ctxtree.go](ctxtree.go):

1. Implement `NewRoot`, `Child`, `Cancel`, `Done`, and `Err` on `*Node`.
2. Cancelling a node cancels its whole subtree exactly once; a node created under an already cancelled parent is born cancelled.
3. Constraint: `-race` clean, `Done()` closes exactly once per node, and no goroutine is required to propagate cancellation.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  cancel the root
Output: every descendant's Done is closed
```

**Example 2:**

```
Input:  cancel a child
Output: siblings stay live
```

**Example 3:**

```
Input:  Child of a cancelled node
Output: already cancelled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cancellation propagation** | The parent walks its children; no polling and no goroutine per node. |
| 2 | **Close-once discipline** | Closing a channel twice panics, so cancellation must be idempotent. |
| 3 | **Lock ordering** | Reused: recursion under locks must not re-enter the same mutex. |

## Hint

Guard with a `cancelled` flag: close `done` only on the first cancel, then recurse into the children.

## Validate

```bash
make verify
```
