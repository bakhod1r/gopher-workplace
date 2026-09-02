# Sorted Unique

## Intuition

Because `cmp.Ordered` implies `comparable`, one constraint buys both the map key and the sort comparison — no second type parameter needed.

## Approach

1. Deduplicate with a `seen` map, preserving first-seen order.
2. Sort the deduplicated slice.
3. Return it.

## Solution

```go
func SortedUnique[T cmp.Ordered](s []T) []T {
	seen := make(map[T]bool, len(s))
	out := make([]T, 0, len(s))
	for _, v := range s {
		if seen[v] {
			continue
		}
		seen[v] = true
		out = append(out, v)
	}
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j] < out[j-1]; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
```

## Walkthrough

`SortedUnique([]int{3, 1, 3})` collects `[3 1]`, skips the repeat, and sorts to `[1 3]`.

## Pitfalls

- Sorting first and then comparing neighbours — valid, but it mutates a copy you forgot to make.
- Using a nested loop for uniqueness, turning the function quadratic.
- Forgetting that the output must be sorted, not first-seen order.
