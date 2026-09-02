# Vector Arithmetic

## Intuition

Rejecting mismatched shapes instead of truncating is a deliberate choice: silently zipping to the shorter length would produce a plausible but wrong score.

## Approach

1. `Add`: return an empty slice on mismatch, otherwise sum element-wise.
2. `Dot`: return zero and `false` on mismatch, otherwise accumulate the products.

## Solution

```go
func Add[T Number](a, b []T) []T {
	out := make([]T, 0, len(a))
	if len(a) != len(b) {
		return make([]T, 0)
	}
	for i := range a {
		out = append(out, a[i]+b[i])
	}
	return out
}

func Dot[T Number](a, b []T) (T, bool) {
	if len(a) != len(b) {
		var zero T
		return zero, false
	}
	var sum T
	for i := range a {
		sum += a[i] * b[i]
	}
	return sum, true
}
```

## Walkthrough

`Dot([]int{1,2}, []int{3,4})` computes `1*3 + 2*4 = 11`.

## Pitfalls

- Zipping to the shorter length and hiding the shape error.
- Accumulating the dot product in `float64` and losing the element type.
- Panicking on a length mismatch.
