# Conversion order in fixed-point

## The idea

Converting a float to an integer **truncates** the fraction. Do it before
scaling and you throw the cents away:

```go
int64(2.50) * 100 // int64(2.50) == 2, then *100 == 200  ✗
int64(2.50 * 100) // 2.50*100 == 250.0, then truncate == 250  ✓
```

Scale in floating point first, convert last.

## Why it matters

Money is best held as integer minor units (cents), but the *conversion* from a
float dollar amount must preserve the fraction. Converting too early is a
classic financial rounding bug that under-charges or under-pays by the cents.

## Watch out

- `int64(x)` rounds toward zero, not to nearest. Add 0.5 (with care for sign) if
  you want nearest.
- Floating dollar inputs are themselves imprecise; prefer integer cents
  end-to-end where possible.
- Multiplying before converting can overflow for huge inputs — bound them.

## Try it yourself

```go
int64(1.99)       // 1
int64(1.99 * 100) // 199
int64(0.07 * 100) // 7 (mind float rounding for some values)
```
