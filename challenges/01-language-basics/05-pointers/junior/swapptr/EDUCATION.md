# Swapping via pointers

## The idea

Dereferencing two pointers and using parallel assignment exchanges the caller's values without a temporary.

## Why it matters

Pointer swap underlies in-place sorts and buffer rotation.

## Watch out

- Swapping the pointers themselves (`a, b = b, a`) does nothing to the caller.
- Dereference to touch the values.
