# Multiple assignment evaluation order

## Intuition

Go evaluates all right-hand-side operands before assigning any left-hand target, enabling rotations and swaps without temporaries.

## Approach

1. Sequential assignment clobbers values before they are used.
2. Use parallel assignment `a, b, c = b, c, a`.

## Solution

```go
func RotateLeft(a, b, c int) (int, int, int) {
	a, b, c = b, c, a
	return a, b, c
}
```

## Walkthrough

The bug's `a = b; b = c; c = a` sets `c` from the already-overwritten `a`. Parallel assignment evaluates the whole right side first, rotating correctly.

## Pitfalls

- `a, b, c = b, c, a` moves all three at once.
- Sequential `a = b; ...; c = a` uses the already-changed `a`.
