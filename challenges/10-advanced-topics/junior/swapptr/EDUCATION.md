# Exchange Two Values Through Pointers

## Intuition

Go evaluates the whole right-hand side before assigning any of it, so a tuple swap needs no temporary and stays correct even when both pointers address the same variable.

## Approach

1. Return early when either pointer is nil.
2. `*a, *b = *b, *a`.

## Solution

```go
// Swap exchanges the values a and b point at.
//
// If either pointer is nil, nothing happens. Nothing is allocated.
//
// Examples:
//
// 	x, y := 1, 2; Swap(&x, &y) => x is 2, y is 1
func Swap(a, b *int) {
	if a == nil || b == nil {
		return
	}
	*a, *b = *b, *a
}
```

## Walkthrough

With `Swap(&x, &x)` the right side reads 5 and 5, then writes 5 and 5 — unchanged, as it should be. A sequential swap with a temporary is also correct; the tuple form is just shorter.

## Pitfalls

- Taking `int` parameters, which swaps copies and does nothing.
- Writing `*a = *b` then `*b = *a`, which loses the original value of `*a`.
