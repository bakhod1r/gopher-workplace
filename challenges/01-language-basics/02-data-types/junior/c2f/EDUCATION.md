# Numeric type conversion

## Intuition

Go never converts numeric types implicitly. To mix an `int` with float math you
convert explicitly, and *when* you convert matters:

```go
float64(c)*9/5 + 32 // float division, keeps 1.8
c*9/5              // int division: 9/5 == 1, fraction lost
```

Converting `c` to `float64` first makes the whole expression floating point.

## Approach

1. Convert c to float64 first so the division is floating point.
2. Apply float64(c)*9/5 + 32.

## Solution

```go
func ToF(c int) float64 {
	return float64(c)*9/5 + 32
}
```

## Walkthrough

ToF(37): float64(37)*9 = 333, /5 = 66.6, +32 = 98.6.

## Pitfalls

- `float64(c) * 9 / 5`: the `9` and `5` become untyped constants that adopt
  float64 here, so the division is floating.
- Converting the *result* (`float64(c*9/5)`) is too late — truncation already
  happened.
- Narrowing conversions (`int(f)`) truncate toward zero and can overflow.
