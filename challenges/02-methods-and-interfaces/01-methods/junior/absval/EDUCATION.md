# Defined Integer Types

## Intuition

`type MyInt int` creates a distinct type. `MyInt(5)` and `int(5)` are different
types — you cannot assign one to the other without conversion. But `MyInt` can
have methods, and those methods can return `MyInt`.

## Approach

1. If `n < 0`, return `-n`.
2. Otherwise, return `n`.

## Solution

```go
func (n MyInt) Abs() MyInt {
	if n < 0 {
		return -n
	}
	return n
}
```

## Walkthrough

For `MyInt(-5)`:
- `-5 < 0` is true.
- `-(-5)` = `5`.

## Pitfalls

- Returning `int` instead of `MyInt` — type mismatch.
- Not handling zero — but zero is already non-negative.
- `math.Abs` works on `float64`, not `int` or `MyInt`.
