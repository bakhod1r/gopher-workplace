# Finalizer Queue

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A handle wraps an OS resource. Callers should close it, but a forgotten handle must not leak the resource forever.

## Task

Implement the stub(s) in [finalizerq.go](finalizerq.go):

1. Implement `NewHandle`, which registers a finalizer as a safety net.
2. Implement `Close`, which releases the resource exactly once and clears the finalizer.
3. Constraint: closing twice is safe, an explicit `Close` must prevent the finalizer from running, and a dropped handle is eventually released by the GC — the test forces collection and checks.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Close then Close
Output: released once
```

**Example 2:**

```
Input:  a handle dropped without Close
Output: the finalizer releases it after a GC
```

**Example 3:**

```
Input:  Close then a GC
Output: still released only once
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **runtime.SetFinalizer** | A last-resort safety net, never the primary release path. |
| 2 | **Finalizer hazards** | They run at an unpredictable time, on a shared goroutine, at most once. |
| 3 | **Explicit release** | Reused: exactly-once cleanup, with the finalizer cleared on the happy path. |

## Hint

Register the finalizer on the handle, and call `runtime.SetFinalizer(h, nil)` inside `Close`.

## Validate

```bash
make verify
```
