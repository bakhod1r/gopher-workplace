# Filter Without A Second Slice

## Intuition

The kept elements can only move left, never right, so the read cursor is always at or ahead of the write cursor. That means you can overwrite as you go without ever losing an element you still need.

## Approach

1. Keep a write index `k` starting at 0.
2. Range over `s`; when the element qualifies, store it at `s[k]` and bump `k`.
3. Return `s[:k]`.

## Solution

```go
// KeepEven returns the even elements of s, in order, reusing s's own
// storage rather than allocating a result.
//
// The elements of s beyond the returned length are unspecified.
//
// Examples:
//
// 	KeepEven([]int{1, 2, 3, 4}) => []int{2, 4}
func KeepEven(s []int) []int {
	k := 0
	for _, v := range s {
		if v%2 == 0 {
			s[k] = v
			k++
		}
	}
	return s[:k]
}
```

## Walkthrough

For [1 2 3 4 6]: 1 is skipped; 2 goes to s[0], k=1; 3 skipped; 4 goes to s[1], k=2; 6 goes to s[2], k=3. Return s[:3] = [2 4 6].

## Pitfalls

- Appending to a fresh slice — correct output, one allocation per call.
- Assuming the tail of `s` is untouched afterwards; it is not.
