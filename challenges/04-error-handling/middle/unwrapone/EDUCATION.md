# Peel One Layer

## Intuition

`Unwrap` is a single step down the chain. `errors.Is` and `errors.As` are the recursive tools built on top of it; using `Unwrap` directly means you want exactly one level.

## Approach

1. Return `errors.Unwrap(err)`.
2. It already returns nil for nil and for non-wrapping errors.

## Solution

```go
return errors.Unwrap(err)
```

## Walkthrough

The two-layer case unwraps to the inner wrapper — still an error with its own cause — not all the way to `ErrBase`.

## Pitfalls

- Looping until the chain ends, which is a different function.
- Assuming every error unwraps; `errors.New` values do not.
- Calling `err.Unwrap()` directly — not every error has that method.
