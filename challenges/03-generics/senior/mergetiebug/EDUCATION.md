# Merge That Loses Stability

## Intuition

`key(b[j]) <= key(a[i])` drains `b` first whenever the keys match, putting the right-hand stream ahead of the left-hand one.

## Approach

1. Take from `b` only when its key is strictly smaller.
2. Otherwise take from `a`.
3. Append whatever remains.

## Solution

```go
func MergeSorted[T any, K cmp.Ordered](a, b []T, key func(T) K) []T {
	out := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if key(b[j]) < key(a[i]) {
			out = append(out, b[j])
			j++
		} else {
			out = append(out, a[i])
			i++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}
```

## Walkthrough

Merging one record keyed 2 from `a` with one keyed 2 from `b` yields the `b` record first.

## Pitfalls

- Testing only with distinct keys, where both operators agree.
- Forgetting one of the two tail appends.
