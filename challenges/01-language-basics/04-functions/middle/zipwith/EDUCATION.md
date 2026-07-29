# Element-wise combination of two sequences

## Intuition

ZipWith iterates two inputs together, bounded by the shorter, applying a two-argument function.

## Approach

1. Iterate up to the shorter length.
2. Combine `a[i]` and `b[i]` with `f`.

## Solution

```go
func ZipWith(a, b []int, f func(int, int) int) []int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, f(a[i], b[i]))
	}
	return out
}
```

## Walkthrough

`ZipWith([1 2 3], [10 20 30], add)` pairs and sums → `[11 22 33]`; mismatched lengths stop at the shorter.

## Pitfalls

- Stop at the shorter length; indexing past it panics.
- Result length equals the shorter input's length.
