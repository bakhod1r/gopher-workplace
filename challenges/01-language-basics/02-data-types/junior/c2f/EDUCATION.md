# Numeric type conversion

## The idea

Go never converts numeric types implicitly. To mix an `int` with float math you
convert explicitly, and *when* you convert matters:

```go
float64(c)*9/5 + 32 // float division, keeps 1.8
c*9/5              // int division: 9/5 == 1, fraction lost
```

Converting `c` to `float64` first makes the whole expression floating point.

## Why it matters

Formulas that look right in math give wrong answers in integer arithmetic
because `/` truncates. Converting at the start preserves precision through the
calculation.

## Watch out

- `float64(c) * 9 / 5`: the `9` and `5` become untyped constants that adopt
  float64 here, so the division is floating.
- Converting the *result* (`float64(c*9/5)`) is too late — truncation already
  happened.
- Narrowing conversions (`int(f)`) truncate toward zero and can overflow.

## Try it yourself

```go
9 / 5            // 1
float64(9) / 5   // 1.8
int(1.99)        // 1
```
