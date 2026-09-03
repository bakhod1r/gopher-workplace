# The Budget Is Per Request, Not Per Call

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A 200ns helper is fast right up until a request calls it ten thousand times, at which point it is two milliseconds and the whole latency budget. Optimisation decisions live at the request level: cost per call times calls per request, against the budget you promised.

## Task

Implement all three in [timerbudget.go](timerbudget.go):

1. `Cost` multiplies ns/op by the calls per request; non-positive inputs cost `0`.
2. `Headroom` returns the budget minus that cost, keeping the sign when it is blown.
3. `Fits` reports whether the cost is within the budget; spending exactly the budget fits.

## Examples

**Example 1:**

```
Input:  Cost(50, 20)
Output: 1000
```

**Example 2:**

```
Input:  Headroom(50, 20, 400)
Output: -600
```

**Example 3:**

```
Input:  Fits(200, 10000, 1000000)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Call count is half the cost** | Per-call latency alone never tells you whether something matters. |
| 2 | **A signed overrun is information** | Clamping to zero throws away how far over you are. |
| 3 | **Inclusive budgets** | "Within 1ms" includes 1ms exactly. |

## Topics used again

Integer arithmetic, guards, boolean returns.

## Hint

`Fits` and `Headroom` should both be expressed in terms of `Cost`.

## Validate

```bash
make verify
```
