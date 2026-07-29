# Saturate to [0,1]

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Colours and probabilities live in [0,1]. Saturating clamps out-of-range values —
and NaN must be pinned to a safe default.

## Task

Implement `Saturate(x)` limiting to [0,1]; NaN → 0.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Saturate(0.5)
Output: 0.5
```

_Explanation:_ already in range

**Example 2:**

```
Input:  Saturate(2)
Output: 1
```

_Explanation:_ above 1 clamps to 1

**Example 3:**

```
Input:  Saturate(-1)
Output: 0
```

_Explanation:_ below 0 clamps to 0

**Example 4:**

```
Input:  Saturate(NaN)
Output: 0
```

_Explanation:_ NaN maps to 0

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Range clamp** | Two comparisons pin the value. |
| 2 | **NaN handling** | NaN fails `<` and `>`; test it explicitly. |
| 3 | **Ordering** | Check NaN (or low bound) first. |

## Hint

Handle NaN first (`math.IsNaN`), then `if x<0 {0}; if x>1 {1}; x`.

## Validate

```bash
make verify
```
