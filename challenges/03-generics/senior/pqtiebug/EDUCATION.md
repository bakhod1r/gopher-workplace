# Insertion That Breaks Stability

## Intuition

Stopping at the first element whose key is not strictly less than `k` puts `v` *before* every equal element, reversing arrival order among equals.

## Approach

1. Scan while the element's key is less than or equal to `k`.
2. Grow the slice by one.
3. Shift the tail right and drop `v` in.

## Solution

```go
func InsertSorted[T any, K cmp.Ordered](s []T, v T, key func(T) K) []T {
	k := key(v)
	i := 0
	for i < len(s) && key(s[i]) <= k {
		i++
	}
	s = append(s, v)
	copy(s[i+1:], s[i:])
	s[i] = v
	return s
}
```

## Walkthrough

Inserting a second priority-2 job into a queue that already holds one lands it at the earlier index, ahead of the job that arrived first.

## Pitfalls

- Assuming ties do not matter for a priority queue — they decide fairness.
- Shifting with a loop that overwrites before it reads; `copy` handles the overlap.
