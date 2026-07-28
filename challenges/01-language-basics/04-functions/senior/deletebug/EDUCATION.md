# The ordered-delete idiom

## The idea

Removing index i keeps the prefix `xs[:i]` and the suffix `xs[i+1:]`; using `xs[i:]` for the suffix keeps the element you meant to drop.

## Why it matters

This exact off-by-one is one of the most-copied Go slice bugs.

## Watch out

- Suffix after deletion starts at `i+1`, not `i`.
- This idiom mutates the backing array; copy first if the caller keeps xs.
