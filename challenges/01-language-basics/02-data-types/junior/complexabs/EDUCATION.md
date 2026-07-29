# Complex numbers

## Intuition

Go has `complex64` and `complex128` as built-in types. Construct with the
built-in `complex(re, im)`; pull the parts with `real(c)` and `imag(c)`, each a
float. The magnitude is `sqrt(re² + im²)`, best computed with `math.Hypot`:

```go
math.Hypot(real(c), imag(c))
```

## Approach

1. Extract parts with the built-ins real(c) and imag(c).
2. Return math.Hypot(re, im), which computes sqrt(re^2+im^2) without overflow.

## Solution

```go
import "math"

func Magnitude(c complex128) float64 {
	return math.Hypot(real(c), imag(c))
}
```

## Walkthrough

Magnitude(complex(3,4)): real=3, imag=4, math.Hypot(3,4)=5.

## Pitfalls

- `real`, `imag`, and `complex` are **built-ins**, not in the `math` package.
- `complex128` holds two `float64`s; `complex64` two `float32`s.
- The `math/cmplx` package has richer operations (`cmplx.Abs`, `cmplx.Phase`).
