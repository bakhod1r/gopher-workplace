# Guarding nil pointers

## Intuition

A nil `*int` holds no address; reading `*p` panics, so a nil check must precede the dereference.

## Approach

1. Guard first: `if p == nil { return def }`.
2. Only after the guard is it safe to `return *p`.
3. A present pointer to zero is different from a missing pointer.

## Solution

```go
func ValueOr(p *int, def int) int {
	if p == nil {
		return def
	}
	return *p
}
```

## Walkthrough

`ValueOr(nil, 5)`: guard catches nil, returns `5`. `ValueOr(&n, 5)` with `n == 9`: guard passes, dereference yields `9`.

## Pitfalls

- `*p` on a nil pointer panics with a nil-dereference.
- Check `p == nil` (or use short-circuit `p != nil && ...`).
