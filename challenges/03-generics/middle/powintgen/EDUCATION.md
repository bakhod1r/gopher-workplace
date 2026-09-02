# Integer Power

## Intuition

The exponent is a count, not a value of the same kind as the base — typing it as `int` keeps the API honest and the loop arithmetic simple.

## Approach

1. Return zero for a negative exponent.
2. Start the result at 1.
3. While the exponent is positive: multiply in on odd, square the base, halve the exponent.

## Solution

```go
func Pow[T Integer](base T, exp int) T {
	if exp < 0 {
		var zero T
		return zero
	}
	var out T = 1
	for exp > 0 {
		if exp%2 == 1 {
			out *= base
		}
		base *= base
		exp /= 2
	}
	return out
}
```

## Walkthrough

`Pow(2, 10)` squares through 2, 4, 16, 256 and multiplies in at the bits set in 10.

## Pitfalls

- Making the exponent a `T`, which forces odd conversions.
- Looping `exp` times, which is exponentially slower for large exponents.
- Returning 0 for `exp == 0`.
