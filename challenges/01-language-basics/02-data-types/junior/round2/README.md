# Round to Decimals

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

`math.Round` rounds to the nearest whole number. To round to N decimals you
scale up, round, and scale back down.

## Task

Implement `Round(x, places)` rounding to `places` decimals, half away from zero.

## Examples

```go
Round(3.14159, 2) // => 3.14
Round(2.5, 0)     // => 3
Round(-2.675, 2)  // => -2.68
```

## Topics to Master

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
