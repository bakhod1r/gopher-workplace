# Abs of the Most-Negative

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`-x` for `int8(-128)` overflows: `128` has no `int8` representation, so `-x`
stays `-128`, and `int(-x)` is `-128`. Two's-complement asymmetry again. Widen
*before* negating.

## Task

Fix the code between the markers in [absval.go](absval.go) so `Abs(-128)` is 128.

## Examples

```go
Abs(-5)   // => 5
Abs(-128) // => 128
Abs(127)  // => 127
```

## Topics to Master

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
