# Sum Of A Range

## Intuition

`(lo+hi)` and `n` are never both odd, so the division is exact; doing it last is what preserves that exactness.

## Approach

1. Return zero for an empty range.
2. Compute the count `hi - lo + 1`.
3. Return `(lo+hi) * n / 2`.

## Solution

```go
func SumRange[T Integer](lo, hi T) T {
	if hi < lo {
		var zero T
		return zero
	}
	n := hi - lo + 1
	return (lo + hi) * n / 2
}
```

## Walkthrough

`SumRange(1, 4)` computes `(1+4) * 4 / 2 = 10`.

## Pitfalls

- Dividing before multiplying, which truncates.
- Looping, which is correct but linear.
- Forgetting the `hi < lo` guard, producing a negative count.
