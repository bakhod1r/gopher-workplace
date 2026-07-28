# Temperature Convert

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 10 min

## Context

A weather widget shows Fahrenheit, but the sensor reports Celsius. Write the
`CToF` helper that converts between them — and watch the constant arithmetic,
where integer division is an easy trap.

## Task

Implement `CToF` in [temperature.go](temperature.go) so it computes
`F = C*(9/5) + 32` correctly. The `freezingF` constant is already provided — do
not change the signature or the tests.

## Examples

```go
CToF(0)   // => 32
CToF(100) // => 212
CToF(-40) // => -40
CToF(37)  // => 98.6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Untyped constants** | `9 / 5` with two integer operands is integer division — the constant is `1`, not `1.8`. |
| 2 | **Constant kind** | Write a fractional constant with a decimal (`9.0 / 5.0` or `1.8`) so it keeps a floating kind. |
| 3 | **Mixed arithmetic** | `float64 * <untyped const>` adopts the constant into `float64`, so the const's *kind* decides the result. |

## Hint

`9 / 5` with two integer operands is *integer* division — the constant truncates
to `1` before it ever touches a float, so `C*1 + 32` runs far too cold. Give the
division a floating-point kind (`9.0 / 5.0`, or just `1.8`).

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.
