# Every Nth

## Intuition

The infinite-loop risk is the real lesson: any strided loop needs its stride validated before the first iteration.

## Approach

1. Return empty for a non-positive stride.
2. Step `i` by `n`, appending each element.

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

`EveryNth([]int{1,2,3,4,5}, 2)` takes indexes 0, 2, and 4.

## Pitfalls

- Omitting the `n <= 0` guard and hanging on `n == 0`.
- Starting at index `n` and dropping the first element.
- Returning nil for an invalid stride.
