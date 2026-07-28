# Untyped constants and precision

## The idea

An untyped constant has no type and carries **arbitrary precision** until it is
assigned to a typed value. Only at that point is it rounded to fit.

```go
const Pi = 3.14159265358979323846  // untyped; ~full float64 precision available
var r float64 = 2
area := Pi * r * r                 // Pi rounds to float64 here
```

Writing extra digits is free: they cost nothing and only round at the point of
use, to whatever float type the expression needs.

## Why it matters

A **typed** constant rounds immediately and permanently:

```go
const P32 float32 = 3.14159265358979 // already truncated to float32 precision
```

Later using `P32` in a `float64` expression cannot recover the lost digits.
Leaving constants untyped keeps the maximum precision available to every caller.

## Watch out

- No type on the left (`const Pi = ...`) keeps it untyped.
- Untyped constants still have a *default* type (float for a float literal) used
  when the context does not force one, e.g. `x := Pi`.
- Precision only matters at conversion time; `Pi * Pi` stays exact until stored.

## Try it yourself

```go
const Third = 1.0 / 3.0 // untyped, high precision
var a float64 = Third   // rounds to float64 here
var b float32 = Third   // rounds to float32 here
```
