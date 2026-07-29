# Variadic with a validity flag

## Intuition

A `...int` plus an `ok bool` cleanly covers the no-argument case without panicking or returning a misleading zero.

## Approach

1. Guard the empty case.
2. Seed `m := nums[0]` and scan the rest.

## Solution

```go
func Max(nums ...int) (int, bool) {
	if len(nums) == 0 {
		return 0, false
	}
	m := nums[0]
	for _, v := range nums[1:] {
		if v > m {
			m = v
		}
	}
	return m, true
}
```

## Walkthrough

`Max(3, 9, 1)` seeds 3, raises to 9; no args returns 0, false.

## Pitfalls

- Seeding from 0 breaks on all-negative inputs; seed from `nums[0]`.
- Guard before indexing.
