# Float64 Exponent Bias

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A numerics library reads the IEEE-754 exponent from the raw bits. The exponent
field is stored **biased by 1023**, but the code subtracts 1024, so every
exponent is off by one.

## Task

Fix the bias between the markers in [floatbits.go](floatbits.go).

## Examples

```go
Exponent(1)    // => 0
Exponent(8)    // => 3
Exponent(0.5)  // => -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **IEEE-754 layout** | 1 sign, 11 exponent, 52 mantissa bits. |
| 2 | **Exponent bias** | Stored value = actual + 1023. |
| 3 | **Bit extraction** | `(bits>>52)&0x7FF`. |

## Hint

`raw - 1023`.

## Validate

```bash
make verify
```
