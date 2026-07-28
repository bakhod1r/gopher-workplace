# Swapping through pointers

## The idea

Reassigning pointer parameters only changes local copies; exchanging the caller's data requires dereferencing both.

## Why it matters

Confusing pointer swap with value swap is a classic no-op bug.

## Watch out

- `a, b = b, a` swaps copies of the addresses.
- `*a, *b = *b, *a` swaps the values the caller holds.
