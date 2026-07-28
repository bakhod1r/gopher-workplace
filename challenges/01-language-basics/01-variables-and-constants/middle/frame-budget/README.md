# Frame Budget

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

A 60 FPS game gets 1/60 s per frame. Computing that from an integer constant
exposes integer-division truncation: `1000000/60 = 16666`, not 16667.

## Task

In [frame.go](frame.go):

1. Implement `FrameBudgetMicros()` = `1_000_000 / TargetFPS` (integer division).
2. Implement `OverBudget(us)` returning true when `us` exceeds the budget.

## Examples

```go
FrameBudgetMicros()      // => 16666
OverBudget(10000)        // => false
OverBudget(20000)        // => true
OverBudget(16666)        // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer const division** | `1000000/60` truncates toward zero → 16666. |
| 2 | **Named constants** | `TargetFPS` documents the magic number. |
| 3 | **Strict vs non-strict** | "over budget" is `>`, not `>=`. |

## Hint

`1_000_000 / TargetFPS` in `int` truncates; don't add a fudge factor.

## Validate

```bash
make verify
```
