# Variadic slice parameters

## The idea

`...[]int` collects trailing slice arguments; spreading each with `g...` builds a flat concatenation.

## Why it matters

Merging config layers or batched results is this pattern.

## Watch out

- A nil group contributes nothing (spreads to zero elements).
- Return empty, not nil-panic, for no groups.
