# The Comparator That Wraps Around

## Intuition

Wrapping the difference in `cmp.Compare` looks safe but the damage is already done: `MaxInt - MinInt` wraps to `-1`, so the comparator claims the largest key is smaller than the smallest.

## Approach

1. Compare the two keys directly with `cmp.Compare`.
2. Never reduce an ordering question to a subtraction on a bounded integer type.

## Solution

```go
func SortByKey[T any](s []T, key func(T) int) {
	slices.SortFunc(s, func(a, b T) int {
		return cmp.Compare(key(a), key(b))
	})
}
```

## Walkthrough

With keys `MaxInt` and `MinInt`, `key(a) - key(b)` evaluates to `-1`, the comparator reports `a < b`, and the two elements are left in their original — wrong — order.

## Pitfalls

- Casting to a wider type to "fix" the subtraction; `int` is already the widest.
- Assuming the tests catch it — ordinary small keys never overflow.
