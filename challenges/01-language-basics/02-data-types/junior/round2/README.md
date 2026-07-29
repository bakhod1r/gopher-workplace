# Round to Decimals

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

`math.Round` rounds to the nearest whole number. To round to N decimals you
scale up, round, and scale back down.

## Task

Implement `Round(x, places)` rounding to `places` decimals, half away from zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Round(3.14159, 2)
Output: 3.14
```

_Explanation:_ Scale by 100, round, divide.

**Example 2:**

```
Input:  Round(2.5, 0)
Output: 3
```

_Explanation:_ math.Round is half away from zero.

**Example 3:**

```
Input:  Round(-2.675, 2)
Output: -2.68
```

_Explanation:_ Half away from zero on negatives too.

**Example 4:**

```
Input:  Round(1.005, 2)
Output: 1.0
```

_Explanation:_ 1.005 is stored slightly below, so it rounds down - a float fact.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **math.Round** | Rounds half away from zero to an integer. |
| 2 | **Scaling** | Multiply by 10^places, round, divide back. |
| 3 | **Float imprecision** | Some decimals are inexact; exact ties are rare. |

## Hint

`p := math.Pow(10, float64(places)); return math.Round(x*p) / p`.

## Validate

```bash
make verify
```
