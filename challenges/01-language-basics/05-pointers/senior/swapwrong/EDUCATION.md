# Swapping through pointers

## Intuition

Reassigning pointer parameters only changes local copies; exchanging the caller's data requires dereferencing both.

## Approach

1. `a, b = b, a` swaps the local **pointer** variables — invisible to the caller.
2. Dereference: `*a, *b = *b, *a` swaps the pointed-at values.

## Solution

```go
func Swap(a, b *int) {
	*a, *b = *b, *a
}
```

## Walkthrough

The buggy line reorders `a` and `b` inside the function only; on return `x` and `y` are unchanged. Writing through `*a`/`*b` swaps the real storage.

## Pitfalls

- `a, b = b, a` swaps copies of the addresses.
- `*a, *b = *b, *a` swaps the values the caller holds.
