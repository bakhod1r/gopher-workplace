# Complex numbers

## The idea

Go has `complex64` and `complex128` as built-in types. Construct with the
built-in `complex(re, im)`; pull the parts with `real(c)` and `imag(c)`, each a
float. The magnitude is `sqrt(re² + im²)`, best computed with `math.Hypot`:

```go
math.Hypot(real(c), imag(c))
```

## Why it matters

Complex arithmetic is native — no library type needed — which is handy for
signal processing and FFTs. `math.Hypot` computes the magnitude without
intermediate overflow or underflow from squaring large components.

## Watch out

- `real`, `imag`, and `complex` are **built-ins**, not in the `math` package.
- `complex128` holds two `float64`s; `complex64` two `float32`s.
- The `math/cmplx` package has richer operations (`cmplx.Abs`, `cmplx.Phase`).

## Try it yourself

```go
c := complex(3, 4)
real(c)            // 3
imag(c)            // 4
math.Hypot(3, 4)   // 5
```
