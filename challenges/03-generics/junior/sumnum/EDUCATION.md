# Sum Numbers

## Intuition

`any` would reject `total += v`. The `Number` union promises every member supports `+`, so the compiler allows the operation for all instantiations at once.

## Approach

1. Declare `var total T`.
2. Add each element.
3. Return the total.

## Solution

```go
func Sum[T Number](s []T) T {
	var total T
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

`Sum([]int{1, 2, 3})` instantiates `T = int`, so `total` starts at `0` and ends at `6`.

## Pitfalls

- Using `[T any]` — `invalid operation: operator + not defined`.
- Starting from a literal `0`, which does not type-check for every `T`.
- Returning `float64` and losing integer precision.
