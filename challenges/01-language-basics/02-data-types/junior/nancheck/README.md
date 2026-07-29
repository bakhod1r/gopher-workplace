# Finite Float Check

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Floating point has special values: `NaN` (not-a-number) and `±Inf`. They arise
from `0.0/0.0`, `math.Sqrt(-1)`, overflow, etc. and must often be rejected.

## Task

Implement `Finite(x)` returning true only for normal finite numbers.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Finite(3.14)
Output: true
```

_Explanation:_ Normal number.

**Example 2:**

```
Input:  Finite(math.NaN())
Output: false
```

_Explanation:_ NaN is not finite.

**Example 3:**

```
Input:  Finite(math.Inf(1))
Output: false
```

_Explanation:_ +Inf is not finite.

**Example 4:**

```
Input:  Finite(math.Inf(-1))
Output: false
```

_Explanation:_ -Inf is not finite.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **NaN** | Not equal to anything, even itself. |
| 2 | **±Inf** | Result of overflow / divide by zero. |
| 3 | **math.IsNaN / IsInf** | The correct way to test these. |

## Hint

`return !math.IsNaN(x) && !math.IsInf(x, 0)`.

## Validate

```bash
make verify
```
