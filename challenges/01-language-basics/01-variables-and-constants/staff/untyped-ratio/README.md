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

## Examples

```go
Value() // => 1.6180555...
```

## Topics to Master

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
