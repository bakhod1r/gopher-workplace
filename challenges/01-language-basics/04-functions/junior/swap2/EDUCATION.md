# Returning a tuple

## Intuition

A function returning `(int, int)` lets callers do `x, y = Swap(x, y)` — Go evaluates the right side fully before assigning.

## Approach

1. Return the arguments in reverse order: `return b, a`.

## Solution

```go
func Swap(a, b int) (int, int) {
	return b, a
}
```

## Walkthrough

`Swap(1, 2)` hands back `2, 1` — multiple return values make this a one-liner.

## Pitfalls

- The function copies its arguments; it cannot mutate the caller's variables.
- Order in the `return` must match the signature.
