# Sorted Copy

## Intuition

Sorting is inherently destructive, so the copy is what makes this function safe. The `less` closure is where `cmp.Ordered` earns its place.

## Approach

1. Allocate `out` with length `len(s)` and `copy` the input into it.
2. Insertion-sort `out` using `<`.
3. Return `out`.

## Solution

```go
func Sorted[T cmp.Ordered](s []T) []T {
	out := make([]T, len(s))
	copy(out, s)
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j] < out[j-1]; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
```

## Walkthrough

`Sorted([]int{3, 1, 2})` copies `[3 1 2]`, sorts the copy to `[1 2 3]`, and leaves the argument untouched.

## Pitfalls

- Sorting `s` directly, which mutates the caller's data.
- Using `make([]T, 0, len(s))` with `copy`, which copies nothing — `copy` is bounded by length, not capacity.
- Comparing indexes instead of elements while sorting.
