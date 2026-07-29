# Untyped constants and precision

## Intuition

An untyped constant has no type and carries **arbitrary precision** until it is
assigned to a typed value. Only at that point is it rounded to fit.

```go
const Pi = 3.14159265358979323846  // untyped; ~full float64 precision available
var r float64 = 2
area := Pi * r * r                 // Pi rounds to float64 here
```

Writing extra digits is free: they cost nothing and only round at the point of
use, to whatever float type the expression needs.

## Approach

1. Declare `Pi` as an untyped constant with 20+ digits.
2. The compiler rounds it to float64 at the point of use.
3. `Area` is `Pi*r*r`.

## Solution

```go
const Pi = 3.14159265358979323846264338327950288

func Area(r float64) float64 {
	return Pi * r * r
}
```

## Walkthrough

`Area(1)` returns Pi rounded to float64; the extra precision avoids compounding rounding error.

## Pitfalls

- No type on the left (`const Pi = ...`) keeps it untyped.
- Untyped constants still have a *default* type (float for a float literal) used
  when the context does not force one, e.g. `x := Pi`.
- Precision only matters at conversion time; `Pi * Pi` stays exact until stored.
