# Divide And Remainder

## Intuition

Go's `%` is a remainder, not a modulus. `-7 % 2` is `-1`, which is why the earlier rotation puzzle needed the double-modulus trick.

## Approach

1. Return zeros and `false` when `b` is zero.
2. Otherwise return `a/b`, `a%b`, and `true`.

## Solution

```go
func DivMod[T Integer](a, b T) (T, T, bool) {
	if b == 0 {
		var zero T
		return zero, zero, false
	}
	return a / b, a % b, true
}
```

## Walkthrough

`DivMod(-7, 2)` returns `-3, -1`: truncation towards zero, with the remainder carrying the dividend's sign.

## Pitfalls

- Letting a zero divisor reach the division and panic.
- Assuming `-7 / 2` rounds down to `-4`.
- Expecting a non-negative remainder for negative dividends.
