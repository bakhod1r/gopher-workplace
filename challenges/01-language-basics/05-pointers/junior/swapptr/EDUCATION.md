# Swapping via pointers

## Intuition

Dereferencing two pointers and using parallel assignment exchanges the caller's values without a temporary.

## Approach

1. `a` and `b` each alias a caller variable.
2. Use parallel assignment `*a, *b = *b, *a`.
3. Go evaluates the whole right side first, so no temporary is needed.

## Solution

```go
func Swap(a, b *int) {
	*a, *b = *b, *a
}
```

## Walkthrough

With `x, y := 1, 2` and `Swap(&x, &y)`:

- The right side `*b, *a` reads `2, 1`.
- Both stores happen at once: `*a = 2`, `*b = 1`.
- Caller now sees `x == 2`, `y == 1`.

## Pitfalls

- Swapping the pointers themselves (`a, b = b, a`) does nothing to the caller.
- Dereference to touch the values.
