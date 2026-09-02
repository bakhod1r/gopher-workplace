# Binary Search

## Intuition

Returning the insertion point makes the function double as "where would this go?", which is what keeps an index sorted as it grows.

## Approach

1. Return `slices.BinarySearch(s, v)` directly.

## Solution

```go
func Find[T cmp.Ordered](s []T, v T) (int, bool) {
	return slices.BinarySearch(s, v)
}
```

## Walkthrough

`Find([]int{1, 3, 5}, 4)` reports `2, false`: inserting `4` at index 2 would keep the slice sorted.

## Pitfalls

- Passing unsorted data and trusting the answer.
- Ignoring the `bool` and treating the index as a hit.
- Reimplementing the search by hand and getting the midpoint arithmetic wrong.
