# Float Tolerance Equality

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

`0.1 + 0.2 != 0.3` in binary floats. Comparing floats needs a tolerance, not
`==`.

## Task

Implement `Equal(a, b, eps)` = true when `|a-b| <= eps`; NaN never equal.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Equal(0.1+0.2, 0.3, 1e-9)
Output: true
```

_Explanation:_ float rounding diff < eps

**Example 2:**

```
Input:  Equal(1.0, 1.001, 1e-9)
Output: false
```

_Explanation:_ difference 0.001 exceeds eps

**Example 3:**

```
Input:  Equal(NaN, NaN, 1)
Output: false
```

_Explanation:_ NaN never equals anything

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Float inexactness** | Most decimals aren't exact in binary. |
| 2 | **Absolute tolerance** | `math.Abs(a-b) <= eps`. |
| 3 | **NaN** | `NaN` fails every comparison; exclude it. |

## Hint

`return math.Abs(a-b) <= eps` — NaN makes `Abs` NaN, and `NaN <= eps` is false,
so NaN is handled naturally.

## Validate

```bash
make verify
```
