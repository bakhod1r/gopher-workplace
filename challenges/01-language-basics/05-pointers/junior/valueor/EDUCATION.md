# Guarding nil pointers

## The idea

A nil `*int` holds no address; reading `*p` panics, so a nil check must precede the dereference.

## Why it matters

Optional pointer arguments require a nil guard to avoid crashes.

## Watch out

- `*p` on a nil pointer panics with a nil-dereference.
- Check `p == nil` (or use short-circuit `p != nil && ...`).
