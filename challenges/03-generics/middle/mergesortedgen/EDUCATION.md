# Merge Sorted

## Intuition

After the main loop one side is exhausted, so appending both remainders is safe: the empty one contributes nothing.

## Approach

1. Walk both slices while each has elements, taking the smaller head.
2. Append the remainder of both slices.

## Solution

```go
func Merge[T cmp.Ordered](a, b []T) []T {
	out := make([]T, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if b[j] < a[i] {
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

`Merge([]int{1}, []int{1})` sees `1 < 1` as false, so it takes from `a` first, then drains `b`.

## Pitfalls

- Concatenating and sorting, which throws away the inputs' order for no reason.
- Using `a[i] < b[j]`, which breaks ties towards `b`.
- Forgetting one of the two remainder appends.
