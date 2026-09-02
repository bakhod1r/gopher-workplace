# Either

## Intuition

Partial inference is not a thing in Go: either every type parameter is inferable from the arguments, or the call must list them all.

## Approach

1. `Left`: set the left slot and the flag.
2. `Right`: set the right slot only.
3. `Unwrap`: return both slots and the flag.

## Solution

```go
func Left[L, R any](v L) Either[L, R] {
	return Either[L, R]{left: v, isLeft: true}
}

func Right[L, R any](v R) Either[L, R] {
	return Either[L, R]{right: v}
}

func (e Either[L, R]) Unwrap() (L, R, bool) {
	return e.left, e.right, e.isLeft
}
```

## Walkthrough

`Right[string, int](5).Unwrap()` returns the zero string, `5`, and `false`.

## Pitfalls

- Calling `Left("e")` and expecting `R` to be inferred.
- Storing a single `any` field and type-asserting later.
- Making the zero value ambiguous by defaulting the flag to true.
