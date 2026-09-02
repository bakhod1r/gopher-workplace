# Generic Swap

## Intuition

`a, b T` means one type parameter covers both arguments — `Swap(1, "x")` will not compile, which is exactly the safety a plain `any` version would lose.

## Approach

1. Return `b, a`.

## Solution

```go
func Swap[T any](a, b T) (T, T) {
	return b, a
}
```

## Walkthrough

`Swap(1, 2)` infers `T = int` from the first argument and checks the second against it, then returns the pair reversed.

## Pitfalls

- Declaring `[T, U any]` when the puzzle wants both values to be the same type.
- Writing a temporary variable — Go's multiple returns make it unnecessary.
- Expecting `Swap` to mutate the caller's variables: it returns new values.
