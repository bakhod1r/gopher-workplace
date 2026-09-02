# Circuit Breaker

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A failing dependency was being hammered by retries. A breaker now opens after repeated failures and lets one probe through after a cooldown.

## Task

Implement the stub(s) in [circuitifc.go](circuitifc.go):

1. Implement `Call` on `*Breaker`: closed passes through; `Threshold` consecutive failures open it; while open, calls fail fast with `ErrOpen`.
2. After `Cooldown` has elapsed, allow exactly one probe — success closes the breaker and resets the count, failure re-opens it.
3. Constraint: while open, the wrapped operation must not be called at all.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Threshold 2; two failures
Output: breaker opens
```

**Example 2:**

```
Input:  a call while open
Output: ErrOpen, the operation is not invoked
```

**Example 3:**

```
Input:  a successful probe after the cooldown
Output: breaker closes
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Circuit breaker** | Fail fast instead of amplifying an outage. |
| 2 | **Clock injection** | Reused: the cooldown is testable without sleeping. |
| 3 | **State machine over an interface** | Reused: the breaker wraps any operation. |

## Hint

Track `openedAt`; while open and within the cooldown, return `ErrOpen` before touching the operation.

## Validate

```bash
make verify
```
