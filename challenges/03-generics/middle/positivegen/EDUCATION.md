# All Positive

## Intuition

The two functions answer the same question at different resolutions: one for a fast gate, one for a diagnostic message.

## Approach

1. `AllPositive`: return `false` at the first element that is not positive.
2. `FirstNonPositive`: return that element's index, or `-1`.

## Solution

```go
func AllPositive[T Signed](s []T) bool {
	for _, v := range s {
		if v <= 0 {
			return false
		}
	}
	return true
}

func FirstNonPositive[T Signed](s []T) int {
	for i, v := range s {
		if v <= 0 {
			return i
		}
	}
	return -1
}
```

## Walkthrough

`AllPositive([]int{1, 0})` rejects at the zero, since zero fails the strict positivity test.

## Pitfalls

- Using `< 0` and accepting zeros.
- Returning `0` from `FirstNonPositive` for "all positive".
- Making `AllPositive` false for an empty slice.
