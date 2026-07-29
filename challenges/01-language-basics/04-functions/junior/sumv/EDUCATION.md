# Variadic functions

## Intuition

The final parameter written `...T` receives a `[]T` built from the trailing call arguments; passing `slice...` forwards an existing slice instead of copying element by element.

## Approach

1. A variadic `nums ...int` arrives as a slice.
2. Range and accumulate from 0.

## Solution

```go
func Sum(nums ...int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}
```

## Walkthrough

`Sum(1,2,3,4)` iterates the four values adding to a total of 10; no arguments yields 0.

## Pitfalls

- Inside the function the parameter is a real slice; a nil/empty one is valid.
- You can pass either loose args OR one spread slice, never both.
