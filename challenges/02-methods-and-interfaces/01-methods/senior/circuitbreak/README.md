# Circuit Breaker

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A failing dependency should not be hammered. A circuit breaker counts
consecutive failures and, once they reach a threshold, opens: further calls fail
fast without touching the dependency at all.

## Task

Implement `Call` on `*Breaker` in [circuitbreak.go](circuitbreak.go):

1. If `IsOpen`, return `errors.New("circuit open")` **without** calling `fn`.
2. Otherwise call `fn()`.
3. On error: increment `ConsecutiveFails`, and set `IsOpen = true` once it
   reaches `Threshold`. Return that error.
4. On success: reset `ConsecutiveFails` to 0 and return nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Threshold 2; Call(failing)
Output: the error; IsOpen == false (1 of 2)
```

**Example 2:**

```
Input:  Call(failing) again
Output: the error; IsOpen == true
```

**Example 3:**

```
Input:  Call(succeeding) while open
Output: error "circuit open"; fn never runs
```

_Explanation:_ fail-fast is the whole point — an open breaker does not consult the dependency.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Guard before side effect** | The open check must precede the call, or the breaker is decoration. |
| 2 | **Consecutive, not total** | A single success resets the counter. |
| 3 | **Boundary comparison** | `>= Threshold`, so a threshold of 2 opens on the second failure. |

## Hint

Check `b.IsOpen` first and return early. Any version that calls `fn` before
looking at the state still hits the failing dependency.

## Validate

```bash
make verify
```
