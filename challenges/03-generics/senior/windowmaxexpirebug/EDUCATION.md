# Sliding Maximum That Never Expires

## Intuition

Without the expiry check the front index can be older than `i-w+1`, so its value is reported for windows it no longer belongs to.

## Approach

1. Evict the front while its index is outside the window.
2. Pop the back while it is no better than the incoming element.
3. Push the index and emit once the first full window is reached.

## Solution

```go
func WindowMax[T cmp.Ordered](s []T, w int) []T {
	out := make([]T, 0)
	if w <= 0 || w > len(s) {
		return out
	}
	dq := make([]int, 0, len(s))
	for i := range s {
		for len(dq) > 0 && dq[0] <= i-w {
			dq = dq[1:]
		}
		for len(dq) > 0 && s[dq[len(dq)-1]] <= s[i] {
			dq = dq[:len(dq)-1]
		}
		dq = append(dq, i)
		if i >= w-1 {
			out = append(out, s[dq[0]])
		}
	}
	return out
}
```

## Walkthrough

In `[5 1 1 1]` with `w = 2`, index 0 stays at the front forever, so every window reports 5.

## Pitfalls

- Evicting after pushing, which can discard the element just added.
- Comparing `dq[0] < i-w`, which keeps one expired index too long.
