# High-Precision Pi

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Untyped constants carry arbitrary precision until assigned to a typed value.
A `Pi` written to 20+ digits rounds correctly to whichever float type uses it.

## Task

In [geo.go](geo.go):

1. Define `Pi` as an untyped constant with ≥20 significant digits.
2. Implement `Area(r)` returning `Pi*r*r`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Area(1)
Output: ~3.14159
```

**Example 2:**

```
Input:  Area(2)
Output: ~12.566
```

**Example 3:**

```
Input:  Area(0)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Untyped constants** | Hold arbitrary precision; no type until used. |
| 2 | **Constant rounding** | Converted to float64 only at the use site. |
| 3 | **Constant vs var** | A typed `var Pi float32` would lose digits immediately. |

## Hint

Leave the type off: `const Pi = 3.14159265358979323846`. The extra digits are
free — they only round when multiplied into a `float64`.

## Validate

```bash
make verify
```
