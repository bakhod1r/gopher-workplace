# When Inference Fails

## Intuition

Type inference in Go flows from argument types only. Languages that infer from the assignment target let you write `var s []int = Empty()`; Go does not, which is why these two functions feel different.

## Approach

1. `Empty`: return `[]T{}`.
2. `ZeroOf`: return `var zero T`.
3. `Wrap`: return `[]T{v}`.

## Solution

```go
func Empty[T any]() []T {
	return []T{}
}

func ZeroOf[T any]() T {
	var zero T
	return zero
}

func Wrap[T any](v T) []T {
	return []T{v}
}
```

## Walkthrough

`var s []int = Empty[int]()` compiles; `Empty()` alone does not, because the parameter cannot be deduced from the left-hand side.

## Pitfalls

- Expecting `var s []int = Empty()` to infer `T` from the variable's type.
- Adding a dummy parameter just to enable inference.
- Returning `nil` from `Empty`, which changes the contract.
