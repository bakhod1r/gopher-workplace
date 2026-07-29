# Pointer equality is identity

## Intuition

`==` on pointers compares the addresses they hold; it is true only when both refer to the same storage.

## Approach

1. Pointer equality compares **addresses**, not pointee values.
2. `return a == b` is true only when both point at the same storage.

## Solution

```go
func Same(a, b *int) bool {
	return a == b
}
```

## Walkthrough

`Same(&x, &x)` compares identical addresses → `true`. `Same(&x, &y)` compares two distinct addresses → `false`, even when `x == y`.

## Pitfalls

- Equal values do NOT imply equal pointers.
- `*a == *b` compares values; `a == b` compares addresses.
