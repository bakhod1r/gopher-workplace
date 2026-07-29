# Returning a selected pointer

## Intuition

Comparing dereferenced values then returning one address hands the caller a live reference to the winner.

## Approach

1. Compare the pointees: `*a` vs `*b`.
2. Return the **pointer** to the larger, not the value.
3. Default to `a` so ties return the first.

## Solution

```go
func MaxPtr(a, b *int) *int {
	if *b > *a {
		return b
	}
	return a
}
```

## Walkthrough

`MaxPtr(&a, &b)` with `*a = 3`, `*b = 8`: `*b > *a` is true, so `b` is returned — the caller can still mutate through it.

## Pitfalls

- Compare `*a`/`*b`, but return `a`/`b`.
- The returned pointer still aliases the caller's variable.
