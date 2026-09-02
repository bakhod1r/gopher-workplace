# Identity

## Intuition

`[T any]` says: this function is written once and compiled for whatever type the caller passes. Inside the body, `T` is an unknown type, so the only safe operation is passing the value around.

## Approach

1. Return `v`.

## Solution

```go
func Identity[T any](v T) T {
	return v
}
```

## Walkthrough

`Identity(7)` infers `T = int` from the argument, so the call type-checks as `Identity[int](7)` and returns an `int` — not an `interface{}` that would need a type assertion.

## Pitfalls

- Writing `func Identity(v any) any` instead — that loses the type, forcing the caller to assert.
- Writing the type argument explicitly (`Identity[int](7)`) when inference already handles it.
- Putting `[T any]` after the parameter list — it goes right after the function name.
