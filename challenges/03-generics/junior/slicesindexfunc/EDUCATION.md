# Index By Predicate

## Intuition

`Index` compares with `==`; `IndexFunc` asks a predicate. The moment your test is anything other than equality, the `Func` variant is the one you want.

## Approach

1. Return `slices.IndexFunc` with a predicate testing `n < 0`.

## Solution

```go
func FirstNegative(nums []int) int {
	return slices.IndexFunc(nums, func(n int) bool { return n < 0 })
}
```

## Walkthrough

`FirstNegative([]int{1, -2, -3})` stops at index 1 — the first match, not the smallest value.

## Pitfalls

- Using `slices.Index`, which cannot express "negative".
- Returning the value instead of the index.
- Returning `0` instead of `-1` when nothing matches.
