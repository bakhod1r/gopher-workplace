# Untyped Constant Division

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`233 / 144` with two integer operands is integer division — the untyped constant
becomes `1`, losing the fraction *before* it ever reaches `float64`. Make at
least one operand a floating literal so the constant division stays exact.

## Task

Fix the single line between the markers in [ratio.go](ratio.go) so `Value()`
returns ≈1.618.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Value()
Output: ~1.618
```

**Example 2:**

```
Input:  233/144 as int
Output: 1 (truncated)
```

**Example 3:**

```
Input:  233.0/144
Output: ~1.618
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer constant division** | `233/144` is `1` when both are integers. |
| 2 | **Untyped float constant** | `233.0/144.0` divides in full precision. |
| 3 | **Type at use** | The constant only rounds when assigned to float64. |

## Hint

`233.0 / 144.0`.

## Validate

```bash
make verify
```
