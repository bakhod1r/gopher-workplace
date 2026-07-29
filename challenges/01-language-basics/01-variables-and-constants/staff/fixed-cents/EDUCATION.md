# Conversion order in fixed-point

## Intuition

Converting a float to an integer **truncates** the fraction. Do it before
scaling and you throw the cents away:

```go
int64(2.50) * 100 // int64(2.50) == 2, then *100 == 200  ✗
int64(2.50 * 100) // 2.50*100 == 250.0, then truncate == 250  ✓
```

Scale in floating point first, convert last.

## Approach

1. `int64(dollars) * 100` truncates dollars before scaling.
2. Scale first: `int64(dollars * 100)`.

## Solution

```go
func Cents(dollars float64) int64 {
	return int64(dollars * 100)
}
```

## Walkthrough

`Cents(0.5)`: the bug truncates 0.5 to 0, then ×100 = 0. Scaling first gives `int64(50.0) = 50`.

## Pitfalls

- `int64(x)` rounds toward zero, not to nearest. Add 0.5 (with care for sign) if
  you want nearest.
- Floating dollar inputs are themselves imprecise; prefer integer cents
  end-to-end where possible.
- Multiplying before converting can overflow for huge inputs — bound them.
