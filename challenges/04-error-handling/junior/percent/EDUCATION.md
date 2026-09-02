# Percentage Of Total

## Intuition

Mixing integer arithmetic with a fractional result is where precision quietly disappears. The conversion has to happen before the division, and the zero denominator has to be rejected before either.

## Approach

1. Reject negative arguments.
2. Reject a zero total.
3. Convert both to `float64`, divide, multiply by 100.

## Solution

```go
if part < 0 || total < 0 {
	return 0, ErrNegative
}
if total == 0 {
	return 0, ErrZeroTotal
}
return float64(part) / float64(total) * 100, nil
```

## Walkthrough

`float64(1) / float64(4) * 100` is `25`. Written as `float64(1/4) * 100` it would be `0`.

## Pitfalls

- Dividing as integers and converting afterwards.
- Checking `total == 0` before the negative guard, so `-4` reports the wrong error.
- Returning `math.Inf` instead of an error for a zero total.
