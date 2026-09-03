# Sort A Small Set Without Reaching For The Heap

## Intuition

A general sorter must work for any length, so it takes an interface and the data escapes. When the length is three and known at compile time, the whole problem collapses into comparisons between registers.

## Approach

1. Order `a` and `b` so `a <= b`.
2. If `b > c`, replace `b` with `c`, then lift it back to `a` if `a` is now larger.
3. Return `b`.

## Solution

```go
// Median3 returns the middle value of a, b and c.
//
// Sorting three ints through a general sorter would box them behind an
// interface. Comparisons alone are enough.
//
// Examples:
//
// 	Median3(3, 1, 2) => 2
func Median3(a, b, c int) int {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b = c
		if a > b {
			b = a
		}
	}
	return b
}
```

## Walkthrough

For (3,1,2): the first swap gives a=1, b=3. b > c so b becomes 2; a (1) is not greater, so the answer is 2.

## Pitfalls

- `a+b+c-min-max` overflows for large ints.
- Missing the equal cases; the median of (5,5,1) is 5, not 1.
