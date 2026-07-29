# Remainder and its sign

## Intuition

`%` gives the remainder of integer division. For parity, `n % 2` is 0 when `n`
is even. The reliable test is `n%2 == 0`.

## Approach

1. Test n%2 == 0 for even (works for negatives).
2. Otherwise return "odd".

## Solution

```go
func Parity(n int) string {
	if n%2 == 0 {
		return "even"
	}
	return "odd"
}
```

## Walkthrough

Parity(-7): -7%2 = -1, not 0, so return "odd".

## Pitfalls

- `a % b` sign follows `a` (the dividend), unlike Python where it follows `b`.
- `%` is only for integers; use `math.Mod` for floats.
- Division by zero panics; `%` by zero panics too.
