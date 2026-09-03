# Sampling That Skips The First

## Intuition

Starting at `n` skips element 0, so the sample is offset by one step and one element short.

## Approach

1. Return early for a non-positive step.
2. Walk from 0 in strides of `n`, collecting each element.

## Solution

```go
func EveryNth[T any](s []T, n int) []T {
	out := make([]T, 0)
	if n <= 0 {
		return out
	}
	for i := 0; i < len(s); i += n {
		out = append(out, s[i])
	}
	return out
}
```

## Walkthrough

`EveryNth([0,1,2,3], 2)` yields `[1 3]` instead of `[0 2]`.

## Pitfalls

- Dropping the `n <= 0` guard, which makes the loop never advance.
- Using `i % n == 0` inside a full scan — correct, but pointlessly slower.
