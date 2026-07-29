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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FrameBudgetMicros()
Output: 16666
```

_Explanation:_ 1_000_000 / 60, integer-divided.

**Example 2:**

```
Input:  OverBudget(20000)
Output: true
```

**Example 3:**

```
Input:  OverBudget(1000)
Output: false
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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
