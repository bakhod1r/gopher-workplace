# Saturation That Never Triggers

## Intuition

The overflow test was copied from signed code. For unsigned types the sum wraps to a small positive number, which the `< 0` test can never catch.

## Approach

1. Add the operands.
2. Compare the sum against one of them.
3. Return `^T(0)` when the sum came out smaller.

## Solution

```go
func SatAdd[T Unsigned](a, b T) T {
	sum := a + b
	if sum < a {
		return ^T(0)
	}
	return sum
}
```

## Walkthrough

`maxUint64 + 1` wraps to `0`, which is not less than zero but is less than `a`.

## Pitfalls

- Porting a signed overflow check to unsigned arithmetic.
- Widening to `uint64` and returning the wrong maximum for narrower types.
- Writing a literal maximum instead of `^T(0)`.
