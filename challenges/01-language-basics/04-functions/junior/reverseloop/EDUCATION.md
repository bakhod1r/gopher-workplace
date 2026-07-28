# Index-based reversal

## The idea

Reading source back-to-front into a new slice keeps the original intact — a copy, not an in-place swap.

## Why it matters

Non-destructive transforms are safer defaults when the caller may still need the input.

## Watch out

- Reversing in place would mutate the caller's slice (shared backing array).
- Start the index at `len(xs)-1`, end at `>= 0`.
