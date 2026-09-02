# The Clamp That Runs Too Late

## Intuition

The clamp compares the sum against limits it can never exceed — the sum has already wrapped into range, so both branches are dead code. The check has to be rearranged so it asks whether the addition *would* overflow, using only values that fit.

## Approach

1. Fetch the limits for `T`.
2. For a positive addend, clamp when `a` exceeds `maxV - b`.
3. For a negative addend, clamp when `a` is below `minV - b`.
4. Otherwise the sum is safe.

## Solution

```go
func AddSat[T Signed](a, b T) T {
	minV, maxV := limits[T]()
	if b > 0 && a > maxV-b {
		return maxV
	}
	if b < 0 && a < minV-b {
		return minV
	}
	return a + b
}

func limits[T Signed]() (T, T) {
	var maxV T
	for x := T(1); x > 0; x <<= 1 {
		maxV |= x
	}
	return -maxV - 1, maxV
}
```

## Walkthrough

`AddSat(int8(100), int8(100))` computes `sum = -56`. That is neither `> 127` nor `< -128`, so it is returned unchanged.

## Pitfalls

- Detecting overflow by widening to `int64`, which cannot work when `T` is already `int64`.
- Writing `maxV - b` without the `b > 0` guard — that expression can itself overflow.
