# Interleave

## Intuition

A single index works because the alternation is positional. Once it runs past either length, exactly one remainder is non-empty.

## Approach

1. Alternate while both slices have an element at `i`.
2. Append `a[i:]` and `b[i:]`.

## Solution

```go
func Interleave[T any](a, b []T) []T {
	out := make([]T, 0, len(a)+len(b))
	i := 0
	for i < len(a) && i < len(b) {
		out = append(out, a[i], b[i])
		i++
	}
	out = append(out, a[i:]...)
	out = append(out, b[i:]...)
	return out
}
```

## Walkthrough

`Interleave([]int{1,2,3}, []int{9})` emits `1, 9`, then drains `a`'s remainder `[2 3]`.

## Pitfalls

- Using two independent indexes that fall out of step.
- Dropping the remainder of the longer slice.
- Starting with `b`, which reverses the intended order.
