# Atomic Implementation Swap

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A running service must switch its rate-limiting strategy without a restart and without a lock on the request path.

## Task

Implement the stub(s) in [atomicswap.go](atomicswap.go):

1. Implement `Set` and `Get` on `*Strategy`, holding an implementation behind `atomic.Pointer`.
2. Implement `Allow`, which dispatches to the current implementation with no lock and no nil panic when nothing is set.
3. Constraint: `-race` clean while swapping implementations under concurrent traffic; a swap must never be observed half-applied.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Allow before any Set
Output: true (fail open)
```

**Example 2:**

```
Input:  Set(DenyAll{}) then Allow
Output: false
```

**Example 3:**

```
Input:  swapping while requests run
Output: each request sees one whole implementation
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Atomic implementation swap** | Hot-swapping behaviour without stopping the world. |
| 2 | **Interfaces in atomic.Pointer** | Store a pointer to a struct holding the interface, not the interface itself. |
| 3 | **Fail-open defaults** | Reused: a nil implementation must not panic on the request path. |

## Hint

`atomic.Pointer[T]` needs a concrete `T`. Wrap the interface in a small struct and store a pointer to that.

## Validate

```bash
make verify
```
