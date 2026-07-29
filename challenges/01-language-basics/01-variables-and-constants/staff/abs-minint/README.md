# Abs of the Most-Negative

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`-x` for `int8(-128)` overflows: `128` has no `int8` representation, so `-x`
stays `-128`, and `int(-x)` is `-128`. Two's-complement asymmetry again. Widen
*before* negating.

## Task

Fix the code between the markers in [absval.go](absval.go) so `Abs(-128)` is 128.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Abs(int8(-5))
Output: 5
```

**Example 2:**

```
Input:  Abs(int8(-128))
Output: 128
```

_Explanation:_ The most negative int8 must widen correctly.

**Example 3:**

```
Input:  Abs(int8(7))
Output: 7
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two's complement asymmetry** | `-MinInt8` overflows int8. |
| 2 | **Widen then negate** | `-int(x)` negates in the wider type. |
| 3 | **Conversion timing** | Convert before the operation that can overflow. |

## Hint

Negate after widening: `return -int(x)` for the negative branch.

## Validate

```bash
make verify
```
