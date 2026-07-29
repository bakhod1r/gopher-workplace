# Storing computed results through a pointer

## Intuition

Calling a pure function without assigning its result is a no-op; the transformed value must be written back through the pointer.

## Approach

1. The bug calls `f(*p)` and drops the result.
2. Store it: `*p = f(*p)`.

## Solution

```go
func Apply(p *int, f func(int) int) {
	*p = f(*p)
}
```

## Walkthrough

`f(*p)` computes the new value but never writes it, so `x` is unchanged. Assigning through `*p` records the transformation.

## Pitfalls

- `f(*p)` returns a value nobody keeps.
- `*p = f(*p)` updates the caller's variable.
