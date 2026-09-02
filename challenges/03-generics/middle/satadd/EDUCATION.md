# Saturating Add

## Intuition

Because unsigned overflow wraps predictably, the post-hoc comparison is a valid detector — a trick that would be unreliable for signed types, where overflow is still wrapping but the comparison is not conclusive.

## Approach

1. Add the two values.
2. Return `^T(0)` when the sum came out below `a`.
3. Otherwise return the sum.

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

Adding 1 to the maximum wraps to 0, which is below `a`, so the guard returns the maximum instead.

## Pitfalls

- Widening to `uint64` and losing the correct maximum for narrower types.
- Writing a literal maximum instead of `^T(0)`.
- Using the same trick on a signed constraint, where it does not hold.
