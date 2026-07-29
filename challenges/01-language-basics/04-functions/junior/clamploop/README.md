# Clamp Each Element

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

Mapping a transform over a slice into a fresh result keeps the caller's data
intact while a `for range` applies the per-element rule.

## Task

Implement `ClampAll` in [clamploop.go](clamploop.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ClampAll([−1 5 99], 0, 10)
Output: [0 5 10]
```

**Example 2:**

```
Input:  ClampAll([3], 0, 10)
Output: [3]
```

**Example 3:**

```
Input:  ClampAll(nil, 0, 10)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Map over a slice** | Build a new slice by appending transformed values. |
| 2 | **Per-element clamp** | Reuse the two-bound clamp logic. |
| 3 | **Non-destructive** | Never write into the input. |

## Hint

Range `xs`; for each, append the clamped value (`if v<lo v=lo; if v>hi v=hi`) to `out`.

## Validate

```bash
make verify
```
