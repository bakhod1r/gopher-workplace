# Sliding Window Maximum

## Intuition

An element smaller than a later one can never be a maximum again, so discarding it immediately is what keeps the deque short and the algorithm linear.

## Approach

1. Guard the invalid window sizes.
2. For each index: drop expired front indexes, drop back indexes whose values are not greater, then append.
3. Once the first full window is reached, emit `s[idx[0]]`.

## Solution

```go
func WindowMax[T cmp.Ordered](s []T, n int) []T {
	out := make([]T, 0)
	if n <= 0 || n > len(s) {
		return out
	}
	idx := make([]int, 0, len(s))
	for i := 0; i < len(s); i++ {
		for len(idx) > 0 && idx[0] <= i-n {
			idx = idx[1:]
		}
		for len(idx) > 0 && s[idx[len(idx)-1]] <= s[i] {
			idx = idx[:len(idx)-1]
		}
		idx = append(idx, i)
		if i >= n-1 {
			out = append(out, s[idx[0]])
		}
	}
	return out
}
```

## Walkthrough

`WindowMax([]int{1,3,2}, 2)` discards index 0 when `3` arrives, so both windows report `3`.

## Pitfalls

- Storing values instead of indexes and being unable to expire them.
- Emitting before the first full window is complete.
- Re-scanning each window, which is correct but quadratic.
