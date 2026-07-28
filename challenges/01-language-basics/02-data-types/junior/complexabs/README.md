# Complex Magnitude

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Go has built-in complex numbers. `real()` and `imag()` extract the parts;
`math.Hypot` gives the magnitude without overflow.

## Task

Implement `Magnitude(c)` returning `sqrt(re^2 + im^2)`.

## Examples

```go
Magnitude(complex(3, 4))   // => 5
Magnitude(complex(0, 0))   // => 0
Magnitude(complex(-3, -4)) // => 5
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **complex128** | Built-in type; `complex(re, im)` constructs it. |
| 2 | **real / imag** | Built-in functions extract the two float64 parts. |
| 3 | **math.Hypot** | Computes sqrt(a²+b²) safely. |

## Hint

`math.Hypot(real(c), imag(c))`.

## Validate

```bash
make verify
```
